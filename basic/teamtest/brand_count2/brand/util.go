package brand

import (
	"reflect"
	"unsafe"
)

func lasIndex(s []byte, c byte) int {
	for i := len(s) - 1; i >= 1; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func lasIndexN(s []byte, n int, c byte) int {
	for i := len(s) - n; i >= 1; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func lasIndexIdx(s []byte, idx int, c byte) int {
	for i := idx - 1; i >= 1; i-- {
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
		v = d - 48
		n *= uint64(10)
		n1 := n + uint64(v)
		n = n1
	}
	return n
}

func parsebyteToInt(b []byte) (n int) {
	lenb := len(b)
	for i := 0; i < lenb; i++ {
		d := b[i]
		v := d - 48
		n *= 10
		n1 := n + int(v)
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

func combinehashBytes(data []byte, xh int) uint64 {
	var h uint64 = 14695981039346656037
	for _, c := range data {
		h = (h ^ uint64(c)) * 1099511628211
	}
	return h + uint64(xh)
}

func Bytes2String(b []byte) (s string) {
	lb := len(b)
	if lb == 0 {
		return ""
	}

	news := make([]byte, lb)
	copy(news, b)
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&news))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

func Date2int(b []byte) int {
	monid := 0
	for i := 4; i < len(b); i++ {
		c := b[i] - 48
		if c == 253 {
			monid = i
		}
	}
	year := parsebyteToInt(b[2:4]) - 7
	month := parsebyteToInt(b[5:monid]) - 1
	day := parsebyteToInt(b[monid+1:])
	count := year*365 + month*30 + day
	return count
}
