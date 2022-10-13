package qrcode_test

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
)

func test() {
	code, err := qrcode.New("https://www.baidu.com", qrcode.Low)
	if err != nil {
		return
	}
	fmt.Println(code.ToString(true))
	code.Level = qrcode.Highest
	code.DisableBorder = false
	png, err := code.PNG(100)
	if err != nil {
		return
	}
	err = os.WriteFile("qrcode.png", png, os.ModeType)
	if err != nil {
		return
	}
	fmt.Println(code.ToString(true))
}
