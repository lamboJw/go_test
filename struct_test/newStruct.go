package struct_test

import "fmt"

func newStruct() {
	lijiawei := newAuthor("李", "家威", 1)
	wangjiawei := newEditor("王家卫")
	book1 := newBook("十万个为什么", "科普", 20, lijiawei, wangjiawei)
	book1.firstName = "123"
	book1.printBook()
	fmt.Println(book1.editor.firstName) // 指明属性中的属性
	fmt.Println(book1.firstName)        // 如结构体内没有同名属性，则自动引用内嵌结构体的属性
	fmt.Println(book1.Author.firstName) // 内嵌结构体的类型名就是它在结构体内的字段名
}
