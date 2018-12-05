package main

type Node struct {
	next     *Node
	playload [64]byte
}

func main() {
	var n Node
	curr := &n
	for {
		curr.next = new(Node)
		curr = curr.next
	}
}
