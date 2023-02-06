package reflect_test

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Enum int

const (
	Zero Enum = iota
	One
	Two
)

func Main() {
	type dog struct {
		num   Enum
		Index int8 `json:"index" id:"100"` //Tag增加json，可以在json编码的时候，指定字段名
	}
	dog1 := dog{Two, 2}
	json_str, _ := json.Marshal(dog1)
	fmt.Println(string(json_str))

	typeOfDog := reflect.TypeOf(dog{Two, 2})
	fmt.Println(typeOfDog.Kind(), typeOfDog.Name())

	if typeOfDog.Kind() == reflect.Struct {
		fieldNums := typeOfDog.NumField()
		for i := 0; i < fieldNums; i++ {
			fmt.Println(typeOfDog.Field(i).Name, typeOfDog.Field(i).Tag)
		}
	}

	var val1 int
	typeOfOne := reflect.ValueOf(&val1) //取地址后，才能修改值，使用.Elem()获取地址对应的变量，然后就可以SetInt()等方法修改
	typeOfOne.Elem().SetInt(666)
	fmt.Println(typeOfOne.Elem().CanSet(), typeOfOne.Elem().CanAddr(), typeOfOne.Elem().Interface(), val1)
}
