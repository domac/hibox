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

//自定义哈希函数
func hashBytes(data []byte) uint64 {
	var h uint64 = 14695981039346656037
	for _, c := range data {
		h = (h ^ uint64(c)) * 1099511628211
	}
	return h
}

//把字符数组转化为无符号整型
func parsebyteToUint64(b []byte) (n uint64) {
	for i := 0; i < len(b); i++ {
		var v byte
		d := b[i]
		v = d - '0'
		n *= uint64(10)
		n1 := n + uint64(v)
		n = n1
	}
	return n
}

func combineArray(a, b []byte) []byte {
	lenA := len(a)
	lenB := len(b)
	slen := lenA + lenB + 1
	sarray := make([]byte, slen)
	copy(sarray[0:lenA], a)
	copy(sarray[lenA+1:], b)
	return sarray
}
