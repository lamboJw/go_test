package list_test

/**
基本实例化：
var node SingleNode
node.data = "123"
node.next = nil

创建结构体的指针类型：
node := new(SingleNode)  //此时返回一个指针类型
node.data = "123"
node.next = nil
在Go语言中，访问结构体指针的成员变量时可以继续使用"."，这是因为Go语言为了方便开发者访问结构体指针的成员变量，
使用了语法糖，将 node.data 形式转换为 (*node).data，即自动将指针类型的调用转为调用指针指向的变量

取结构体的地址来实例化：
node := &SingleNode{}	//此时返回的是这个结构体的地址。对结构体进行“&”取地址操作时，视为对其进行一次new的实例化操作，然后返回地址
node.data = "123"
node.next = nil
也可以取地址的同时赋值：
node := &SingleNode{
	data: "123",
	next: nil
}
*/

import "fmt"

type SingleNode struct {
	data interface{}
	next *SingleNode
}

func (n *SingleNode) GetData() interface{} {
	return n.data
}

type SingleList struct {
	length   int
	headNode *SingleNode
	tailNode *SingleNode
}

func (list *SingleList) Append(data string) {
	node := &SingleNode{
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

func (list *SingleList) Pop() NodeInterface {
	if list.length == 0 {
		return nil
	}
	returnNode := list.tailNode
	lastNode := list.headNode
	var lastSecondNode *SingleNode = nil
	for lastNode.next != nil {
		lastSecondNode = lastNode
		lastNode = lastNode.next
	}
	if lastSecondNode == nil {
		list.headNode = nil
	} else {
		lastSecondNode.next = nil
	}
	list.tailNode = lastSecondNode
	list.length -= 1
	return returnNode
}

func (list *SingleList) Unshift(data string) {
	node := &SingleNode{
		data: data,
		next: list.headNode,
	}
	list.headNode = node
	if list.tailNode == nil {
		list.tailNode = node
	}
	list.length += 1
}

func (list *SingleList) Shift() NodeInterface {
	if list.length == 0 {
		return nil
	}
	returnNode := list.headNode
	list.headNode = returnNode.next
	if returnNode.next == nil { // 弹出的结点的下一结点为空，说明链表已经弹出所有节点
		list.tailNode = nil // 需要把尾节点也标识为空
	}
	list.length -= 1
	return returnNode
}

func (list *SingleList) Traverse() {
	fmt.Println("链表长度：", list.length)
	p := list.headNode
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func newSingleList() *SingleList {
	return &SingleList{
		length:   0,
		headNode: nil,
		tailNode: nil,
	}
}
