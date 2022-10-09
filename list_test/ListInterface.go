package list_test

type ListInterface interface {
	Append(data string)
	Pop() NodeInterface
	Unshift(data string)
	Shift() NodeInterface
	Traverse()
}
