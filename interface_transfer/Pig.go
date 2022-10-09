package interface_transfer

import "fmt"

type Pig struct {
}

func (p *Pig) walk() {
	fmt.Println("猪在走")
}

//func (p *Pig) fly() {
//	fmt.Println("猪不会飞")
//}
