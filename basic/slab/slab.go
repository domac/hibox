package slab

type Pool interface {
	Alloc(int) []byte
	Free([]byte)
}
