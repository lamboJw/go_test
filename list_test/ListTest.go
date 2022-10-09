package list_test

func ListTest() {
	nodeType := "double"
	list := GetList(nodeType)
	list.Append("1")
	list.Unshift("2")
	list.Unshift("3")
	list.Append("4")
	list.Traverse()
	list.Pop()
	list.Shift()
	list.Traverse()
}

func GetList(nodeType string) ListInterface {
	var list ListInterface = nil
	if nodeType == "single" {
		list = newSingleList()
	} else {
		list = newDoubleList()
	}
	return list
}
