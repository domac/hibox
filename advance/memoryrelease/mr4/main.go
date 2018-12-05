package main

type Node struct {
	next     *Node
	playload [64]byte
}

//用二级指针，new(*Node)会被逃逸优化，但new(Node)就不会了
func main() {
	curr := new(*Node)
	*curr = new(Node)
	for {
		(*curr).next = new(Node)
		*curr = (*curr).next
	}
}

//用return强制逃逸，避免优化
func f() *Node {
	curr := new(Node)
	for {
		curr.next = new(Node)
		curr = curr.next
	}
	return curr
}

// func main() {
// 	f()
// }
