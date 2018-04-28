package main

import (
	"flag"
	"fmt"
	"go/types"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"os"
	"path/filepath"
	"strings"
)

//涉及查询方法结构
type QueryMethod struct {
	Func     *types.Func
	SSA      *ssa.Function
	ArgCount int
	Param    int
}

type sqlPackage struct {
	packageName string
	paramNames  []string
	enable      bool
}

var sqlPackages = []sqlPackage{
	{
		packageName: "database/sql",
		paramNames:  []string{"query"},
	},
}

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "show detail")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-v] package1 [package 2]...\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	pkgs := flag.Args()
	if len(pkgs) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	fmt.Println("欢迎使用SQL注入分析工具")

	c := loader.Config{}

	for _, pkg := range pkgs {
		pkg, _ = filepath.Abs(pkg)
		if strings.Count(pkg, "src") >= 2 {
			srcIndex := strings.Index(pkg, "src") + 4
			pkg = pkg[srcIndex:]
		}
		c.Import(pkg)
	}

	//根据目标路径目录，加载目录下的文件，生产“加载器程序”
	loaderProg, err := c.Load()

	if err != nil {
		fmt.Printf("加载包异常 %v: %v \n", pkgs, err)
		os.Exit(2)
	}

	//整理import的所有包路径（文件头的import信息）
	imports := func(p *loader.Program) map[string]interface{} {
		pkgs := make(map[string]interface{})
		for _, pkg := range p.AllPackages {
			if pkg.Importable {
				pkgs[pkg.Pkg.Path()] = nil
			}
		}
		return pkgs
	}(loaderProg)

	//查找是否包含sql库
	existOne := false
	for i := range sqlPackages {
		//是否存在目标分析的包
		if _, exists := imports[sqlPackages[i].packageName]; exists {
			if verbose {
				fmt.Printf("待检测库: %s\n", sqlPackages[i].packageName)
			}
			sqlPackages[i].enable = true
			existOne = true
		}
	}
	if !existOne {
		fmt.Printf("在%v的包中没有含有所支持的database driver", pkgs)
		os.Exit(2)
	}

	//构建SSA程序
	ssaProg := ssautil.CreateProgram(loaderProg, 0)
	ssaProg.Build()

	qms := make([]*QueryMethod, 0)

	for i := range sqlPackages {
		sp := sqlPackages[i]
		if sp.enable {
			qms = append(qms, findQueryMethods(sp, loaderProg.Package(sp.packageName).Pkg, ssaProg)...)
		}
	}

	if verbose {
		fmt.Println("database driver 待检测函数:")
		for _, m := range qms {
			fmt.Printf(" ~ %s (目标参数位置 %d)\n", m.Func, m.Param)
		}
		fmt.Println("")
	}

	//获取包含main方法的包
	mains := func(prog *loader.Program, ssaProg *ssa.Program) []*ssa.Package {
		ips := prog.InitialPackages()
		mains := make([]*ssa.Package, 0, len(ips))
		for _, info := range ips {
			ssaPkg := ssaProg.Package(info.Pkg)
			if ssaPkg.Func("main") != nil {
				mains = append(mains, ssaPkg)
			}
		}
		return mains
	}(loaderProg, ssaProg)

	if len(mains) == 0 {
		fmt.Println("没有找到相关main方法")
		os.Exit(2)
	}

	//指向分析
	res, err := pointer.Analyze(&pointer.Config{
		Mains:          mains,
		BuildCallGraph: true,
	})

	if err != nil {
		fmt.Printf("执行pointer分析异常: %v\n", err)
		os.Exit(2)
	}

	riskCalls := FindNonConstCalls(res.CallGraph, qms)

	if len(riskCalls) == 0 {
		println("没有发现注入漏洞")
		return
	}

	fmt.Printf("发现 %d 个潜在的注入风险:\n", len(riskCalls))

	showMap := make(map[string][]ssa.CallInstruction)
	for _, ci := range riskCalls {

		pos := loaderProg.Fset.Position(ci.Pos())

		cis := showMap[pos.Filename]
		cis = append(cis, ci)
		showMap[pos.Filename] = cis
	}

	for fileName, cis := range showMap {
		dir := filepath.Dir(fileName)
		fmt.Printf("\n%s", dir)
		for _, ci := range cis {
			pos := loaderProg.Fset.Position(ci.Pos())
			c := strings.Replace(pos.String(), dir, "", 1)
			fmt.Printf("\n+ %s", c[1:])
		}
		println()
	}

	return

}

func FindNonConstCalls(cg *callgraph.Graph, qms []*QueryMethod) []ssa.CallInstruction {
	cg.DeleteSyntheticNodes()

	okFuncs := make(map[*ssa.Function]struct{}, len(qms))

	for _, m := range qms {
		okFuncs[m.SSA] = struct{}{}
	}

	riskCalls := make([]ssa.CallInstruction, 0)

	for _, m := range qms {
		node := cg.CreateNode(m.SSA)
		for _, edge := range node.In {
			if _, ok := okFuncs[edge.Site.Parent()]; ok {
				continue
			}
			isInternalSQLPkg := false

			for _, pkg := range sqlPackages {
				if pkg.packageName == edge.Caller.Func.Pkg.Pkg.Path() {
					isInternalSQLPkg = true
					break
				}
			}

			if isInternalSQLPkg {
				continue
			}

			cc := edge.Site.Common()
			args := cc.Args
			if len(args) == m.ArgCount+1 {
				args = args[1:]
			} else if len(args) != m.ArgCount {
				panic("arg count mismatch")
			}
			v := args[m.Param]

			if _, ok := v.(*ssa.Const); !ok {
				if inter, ok := v.(*ssa.MakeInterface); ok && types.IsInterface(v.(*ssa.MakeInterface).Type()) {
					if inter.X.Referrers() == nil || inter.X.Type() != types.Typ[types.String] {
						continue
					}
				}

				riskCalls = append(riskCalls, edge.Site)
			}
		}
	}

	return riskCalls
}

//针对 database/sql这类型的包查目标检测函数
func findQueryMethods(sp sqlPackage, targetPkg *types.Package, ssaProg *ssa.Program) []*QueryMethod {
	methods := make([]*QueryMethod, 0)
	//pkgScope为最原始的包范围
	pkgRootScope := targetPkg.Scope()
	for _, name := range pkgRootScope.Names() {
		methodObj := pkgRootScope.Lookup(name)
		if !methodObj.Exported() {
			continue
		}

		if _, ok := methodObj.(*types.TypeName); !ok {
			continue
		}

		//获取获取嵌套改method的方法
		methodName := methodObj.Type().(*types.Named)
		for i := 0; i < methodName.NumMethods(); i++ {
			m := methodName.Method(i)
			if !m.Exported() {
				continue
			}
			//println("------", m.String())
			s := m.Type().(*types.Signature)
			//println(s.String())
			if num, ok := FuncHasQuery(sp, s); ok {
				methods = append(methods, &QueryMethod{
					Func:     m,
					SSA:      ssaProg.FuncValue(m),
					ArgCount: s.Params().Len(),
					Param:    num,
				})
			}
		}
	}
	return methods
}

func FuncHasQuery(sp sqlPackage, s *types.Signature) (offset int, ok bool) {
	params := s.Params()
	for i := 0; i < params.Len(); i++ {
		v := params.At(i)
		for _, paramName := range sp.paramNames {
			if v.Name() == paramName {
				return i, true
			}
		}
	}
	return 0, false
}
