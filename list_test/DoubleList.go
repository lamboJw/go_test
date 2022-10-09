package list_test

import "fmt"

type DoubleNode struct {
	prev *DoubleNode
	data interface{}
	next *DoubleNode
}

func (n *DoubleNode) GetData() interface{} {
	return n.data
}

type DoubleList struct {
	length   int
	headNode *DoubleNode
	tailNode *DoubleNode
}

func (list *DoubleList) Append(data string) {
	node := &DoubleNode{
		prev: list.tailNode,
		data: data,
		next: nil,
	}
	if list.headNode == nil {
		list.headNode = node
	} else {
		list.tailNode.next = node
	}
	list.tailNode = node
	list.length += 1
}

func (list *DoubleList) Pop() NodeInterface {
	returnNode := list.tailNode
	if returnNode == nil {
		return nil
	}
	list.tailNode = list.tailNode.prev // 用弹出尾结点的前一个节点作为当前尾结点
	if list.tailNode == nil {          // 如果当前尾结点为空，说明已弹出所有结点
		list.headNode = nil
	} else {
		list.tailNode.next = nil // 还有剩余结点，需要把尾结点的下一个节点设置为空，这样才不会保留弹出结点
	}
	list.length -= 1
	return returnNode
}

func (list *DoubleList) Unshift(data string) {
	node := &DoubleNode{
		prev: nil,
		data: data,
		next: list.headNode,
	}
	if list.headNode != nil {
		list.headNode.prev = node
	}
	list.headNode = node
	if list.tailNode == nil {
		list.tailNode = node
	}
	list.length += 1
}

func (list *DoubleList) Shift() NodeInterface {
	if list.headNode == nil {
		return nil
	}
	returnNode := list.headNode
	list.headNode = returnNode.next
	if list.headNode != nil { // 新的头节点不为空，则把头结点的前节点地址设置为空，去除与弹出结点的关联
		list.headNode.prev = nil
	} else {
		list.tailNode = nil // 弹出结点的下一结点为空，则说明已弹出所有结点，需要把尾结点设置为空
	}
	list.length -= 1
	return returnNode
}

func (list *DoubleList) Traverse() {
	fmt.Println("链表长度：", list.length)
	p := list.headNode
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func newDoubleList() *DoubleList {
	return &DoubleList{
		length:   0,
		headNode: nil,
		tailNode: nil,
	}
}
