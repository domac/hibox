package scan

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"path/filepath"
	"strings"
)

//涉及查询方法结构
type QueryFunctionModel struct {
	Func     *types.Func
	SSA      *ssa.Function
	ArgCount int
	Param    int
}

func DoScan(pkgs []string) []string {

	result := []string{}

	c := loader.Config{}

	for _, pkg := range pkgs {
		if pkg[0] == '/' || pkg[0] == '~' {
			pkg, _ = filepath.Abs(pkg)
			if strings.Count(pkg, "src") >= 1 {
				srcIndex := strings.Index(pkg, "src") + 4
				pkg = pkg[srcIndex:]
			}
		}
		c.Import(pkg)
	}

	//根据目标路径目录，加载目录下的文件，生产“加载器程序”
	loaderProg, err := c.Load()

	if err != nil {
		s := fmt.Sprintf("加载包异常 %v: %v", pkgs, err)
		result = append(result, s)
		return result
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
	for i := range scanPckagesList {
		//是否存在目标分析的包
		if _, exists := imports[scanPckagesList[i].packageName]; exists {
			s := fmt.Sprintf("待检测库: %s", scanPckagesList[i].packageName)
			result = append(result, s)
			scanPckagesList[i].enable = true
			existOne = true
		}
	}
	if !existOne {
		s := fmt.Sprintf("在%v的包中没有含有所支持的database driver", pkgs)
		result = append(result, s)
		return result
	}

	//构建SSA程序
	ssaProg := ssautil.CreateProgram(loaderProg, 0)
	ssaProg.Build()

	qms := make([]*QueryFunctionModel, 0)

	for i := range scanPckagesList {
		sp := scanPckagesList[i]
		if sp.enable {
			qms = append(qms, findQueryMethods(sp, loaderProg.Package(sp.packageName).Pkg, ssaProg)...)
		}
	}

	//if verbose {
	s := fmt.Sprintln("database driver 待检测函数:")
	result = append(result, s)
	for _, m := range qms {
		fmt.Printf(" ~ %s (目标参数位置 %d)\n", m.Func, m.Param)
	}
	result = append(result, "")
	//}

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
		result = append(result, "没有找到相关main方法")
		return result
	}

	//指向分析
	res, err := pointer.Analyze(&pointer.Config{
		Mains:          mains,
		BuildCallGraph: true,
	})

	if err != nil {
		s := fmt.Sprintf("执行pointer分析异常: %v", err)
		result = append(result, s)
		return result
	}

	riskCalls := FindNonConstCalls(res.CallGraph, qms)

	if len(riskCalls) == 0 {
		println("没有发现注入漏洞")
		return result
	}

	result = append(result, fmt.Sprintf("发现 %d 个潜在的注入风险:", len(riskCalls)))

	showMap := make(map[string][]ssa.CallInstruction)
	for _, ci := range riskCalls {

		pos := loaderProg.Fset.Position(ci.Pos())

		cis := showMap[pos.Filename]
		cis = append(cis, ci)
		showMap[pos.Filename] = cis
	}

	for fileName, cis := range showMap {
		dir := filepath.Dir(fileName)
		result = append(result, fmt.Sprintf("%s", dir))
		for _, ci := range cis {
			pos := loaderProg.Fset.Position(ci.Pos())
			c := strings.Replace(pos.String(), dir, "", 1)
			result = append(result, fmt.Sprintf("+ %s", c[1:]))
		}
		result = append(result, "")
	}

	return result
}

func FindNonConstCalls(cg *callgraph.Graph, qms []*QueryFunctionModel) []ssa.CallInstruction {
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

			for _, pkg := range scanPckagesList {
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
func findQueryMethods(sp scanPackage, targetPkg *types.Package, ssaProg *ssa.Program) []*QueryFunctionModel {
	methods := make([]*QueryFunctionModel, 0)
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
			s := m.Type().(*types.Signature)
			if num, ok := FuncHasQuery(sp, s); ok {
				methods = append(methods, &QueryFunctionModel{
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

func FuncHasQuery(sp scanPackage, s *types.Signature) (offset int, ok bool) {
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

type scanPackage struct {
	packageName string
	paramNames  []string
	enable      bool
}

var scanPckagesList = []scanPackage{
	{
		packageName: "database/sql",
		paramNames:  []string{"query"},
	},
}
