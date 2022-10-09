package struct_test

import "fmt"

type Book struct {
	name      string
	price     float32
	subject   string
	firstName string  // 如果结构体内包含内嵌结构体的同名属性，调用时不指明类型名，则默认使用本结构体的属性
	*Author           // 类似继承，Book的实例可以直接调用Author的属性和方法
	editor    *Editor // 单纯的属性，需要通过属性名调用Editor的属性和方法
}

func (book *Book) printBook() {
	fmt.Println("书名：", book.name)
	fmt.Println("类别：", book.subject)
	fmt.Println("价格：", book.price)
	fmt.Println("作者：")
	book.printAuthor()
}

func newBook(name string, subject string, price float32, author *Author, editor *Editor) *Book {
	return &Book{
		name:    name,
		price:   price,
		subject: subject,
		Author:  author, //内嵌结构体的字段名，就是它的类型名
		editor:  editor,
	}
}
