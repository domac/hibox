package brand

const FIELDS_IDX = 4

//自定义反向分隔处理
//s[0]: 日期
//s[1]: 销售额
//s[2]: 仓库
//s[3]: 描述
//s[4]: 品牌
func genSpaceSplit(s []byte) [][]byte {
	//预分配数组
	a := make([][]byte, FIELDS_IDX+2)
	i := 0
	for i < FIELDS_IDX {
		m := lasIndex(s, ' ')
		if m < 0 {
			break
		}
		a[i] = s[m+1:]
		s = s[:m]
		i++
	}
	a[i] = s
	return a[:i+1]
}

func lasIndex(s []byte, c byte) int {
	for i := len(s) - 1; i >= 1; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}
