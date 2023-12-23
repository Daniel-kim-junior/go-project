package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func Append[T any](rNode *Node[T], lNode *Node[T]) *Node[T] {
	rNode.next = lNode
	return lNode
}

func main() {
	root := &Node[int]{nil, 10}
	node := root

	for n := 2; n < 4; n++ {
		node = Append(node, &Node[int]{nil, n * 10})
	}

	for n := root; n != nil; n = n.next {
		fmt.Printf("val %d\n", n.val)
	}
	fmt.Println()
	fmt.Println()

	tmp := root.next.next
	Append(root.next, &Node[int]{tmp, 100})
	for n := root; n != nil; n = n.next {
		fmt.Printf("val %d\n", n.val)
	}

}
