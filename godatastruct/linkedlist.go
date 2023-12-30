package godatastruct

type Item interface {
}

type Node struct {
	Data     Item
	nextNode *Node
	preNode  *Node
}

type LinkedList interface {
	Len() int
	PushBack(item Item)
	PushFront(itme Item)
	PopBack() Item
	PopFront() Item
}

type DoubleLinkedList struct {
	rootNode *Node
	tailNode *Node
	len      int
}

func (d *DoubleLinkedList) PushBack(node Item) {
	d.len++
	newNode := &Node{node, nil, nil}
	if d.rootNode == nil {
		d.rootNode = newNode
		d.tailNode = newNode
		return
	}
	d.tailNode.nextNode, newNode.preNode, d.tailNode = newNode, d.tailNode, newNode
}

func (d *DoubleLinkedList) PushFront(node Item) {
	d.len++
	newNode := &Node{node, nil, nil}
	if d.rootNode == nil {
		d.rootNode = newNode
		d.tailNode = newNode
		return
	}
	d.rootNode.preNode, newNode.nextNode, d.rootNode = newNode, d.rootNode, newNode
}

func (d *DoubleLinkedList) PopFront() *Node {
	if d.len == 0 {
		return nil
	}
	d.len--
	popNode := d.rootNode
	d.rootNode, d.rootNode.nextNode.preNode = d.rootNode.nextNode, nil
	return popNode
}

func (d *DoubleLinkedList) PopBack() *Node {
	if d.len == 0 {
		return nil
	}
	d.len--
	popNode := d.tailNode
	d.tailNode, d.tailNode.nextNode = d.tailNode.preNode, nil
	return popNode
}

func (d *DoubleLinkedList) Len() int {
	return d.len
}
