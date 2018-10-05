package slab

type ChanPool struct {
	classes []chanClass
	minSize int
	maxSize int
}

func NewChanPool(minSize, maxSize, factor, pageSize int) *ChanPool {
	pool := &ChanPool{
		classes: make([]chanClass, 0, 10),
		minSize: minSize,
		maxSize: maxSize,
	}

	for chunkSize := minSize; chunkSize <= maxSize && chunkSize <= pageSize; chunkSize *= factor {
		chunkLen := pageSize / chunkSize

		c := chanClass{
			size:   chunkSize,
			page:   make([]byte, pageSize),
			chunks: make(chan []byte, chunkLen),
		}
		for i := 0; i < chunkLen; i++ {
			mem := c.page[i*chunkSize : (i+1)*chunkSize : (i+1)*chunkSize]
			c.chunks <- mem

		}
		pool.classes = append(pool.classes, c)
	}
	return pool
}

func (pool *ChanPool) Alloc(size int) []byte {
	if size <= pool.maxSize {
		for i := 0; i < len(pool.classes); i++ {
			if pool.classes[i].size >= size {
				mem := pool.classes[i].Pop()
				if mem != nil {
					return mem[:size]
				}
				break
			}
		}
	}
	return make([]byte, size)
}

func (pool *ChanPool) Free(mem []byte) {
	if size := cap(mem); size <= pool.maxSize {
		for i := 0; i < len(pool.classes); i++ {
			if pool.classes[i].size == size {
				pool.classes[i].Push(mem)
				break
			}
		}
	}
}

type chanClass struct {
	size   int
	page   []byte
	chunks chan []byte
}

func (c *chanClass) Push(mem []byte) {
	select {
	case c.chunks <- mem:
	default:
		mem = nil
	}
}

func (c *chanClass) Pop() []byte {
	select {
	case mem := <-c.chunks:
		return mem
	default:
		return nil
	}
}
