package godatastruct

type Item interface {
}

type Node struct {
	Data     Item
	NextNode *Node
	PreNode  *Node
}

type LinkedList interface {
	Len() int
	PushBack(item Item)
	PushFront(itme Item)
	PopBack() Item
	PopFront() Item
}

type DoubleLinkedList struct {
	RootNode *Node
	TailNode *Node
	len      int
}

func (d *DoubleLinkedList) PushBack(node Item) {
	d.len++
	newNode := &Node{node, nil, nil}
	if d.RootNode == nil {
		d.RootNode = newNode
		d.TailNode = newNode
		return
	}
	d.TailNode.NextNode, newNode.PreNode, d.TailNode = newNode, d.TailNode, newNode
}

func (d *DoubleLinkedList) PushFront(node Item) {
	d.len++
	newNode := &Node{node, nil, nil}
	if d.RootNode == nil {
		d.RootNode = newNode
		d.TailNode = newNode
		return
	}
	d.RootNode.PreNode, newNode.NextNode, d.RootNode = newNode, d.RootNode, newNode
}

func (d *DoubleLinkedList) PopFront() *Node {
	if d.len == 0 {
		return nil
	}
	d.len--
	popNode := d.RootNode
	d.RootNode, d.RootNode.NextNode.PreNode = d.RootNode.NextNode, nil
	return popNode
}

func (d *DoubleLinkedList) PopBack() *Node {
	if d.len == 0 {
		return nil
	}
	d.len--
	popNode := d.TailNode
	d.TailNode, d.TailNode.NextNode = d.TailNode.PreNode, nil
	return popNode
}

func (d *DoubleLinkedList) Len() int {
	return d.len
}
