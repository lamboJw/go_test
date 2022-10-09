package struct_test

import "fmt"

type Author struct {
	firstName  string
	secondName string
	sex        int8
}

func (author *Author) printAuthor() {
	fmt.Println("姓名：", author.firstName+author.secondName)
	var sex string
	switch author.sex {
	case 1:
		sex = "男"
	case 2:
		sex = "女"
	default:
		sex = "未知"
	}
	fmt.Println("性别：", sex)
}

func newAuthor(firstName string, secondName string, sex int8) *Author {
	return &Author{
		firstName:  firstName,
		secondName: secondName,
		sex:        sex,
	}
}
