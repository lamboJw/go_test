package interface_transfer

import "fmt"

type Bird struct {
}

func (b *Bird) fly() {
	fmt.Println("鸟在飞")
}

func (b *Bird) walk() {
	fmt.Println("鸟在走")
}