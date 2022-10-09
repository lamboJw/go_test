package interface_transfer

import (
	"fmt"
	"go_test/list_test"
	"reflect"
)

func Main() {
	var bird interface{} = new(Bird)
	f, isFlyer := bird.(Flyer) // 将bird类型转换为Flyer接口并保存到f中
	if isFlyer {
		f.fly()
	}
	w, isWalker := bird.(Walker) // 将bird类型转换为walker接口并保存到w中，不能保存到f中，因为f是Flyer类型，Walker不存在fly()方法，所以类型不对
	if isWalker {
		w.walk()
	}

	a, err := f.(*Bird) // 将一个Flyer接口的对象转换回Bird类型，然后调用Bird的其他方法
	if err {
		a.walk()
	}
	fmt.Printf("f=%p, w=%p, a=%p \n", f, w, a) // 类型转来转去，指针还是相同的

	p1 := new(Pig)
	fmt.Println("p1的类型", reflect.TypeOf(p1))
	var w2 Walker = p1
	fmt.Println("w2的类型", reflect.TypeOf(w2))
	p2, p2Err := w2.(*Bird) // 虽然w2是Walker类型，Bird也实现了Walker接口，按理说是可以转换的，但是w2存储的实例是Pig类型，所以不能转换
	fmt.Println("p2的类型", reflect.TypeOf(p2), "转换结果：", p2Err)

	var a1 interface{} = 100
	var b interface{} = "100"
	c, bErr := b.(int)
	fmt.Println("转换结果", bErr, c, a1 == c)

	// type-switch，类型断言的加强版
	var w3 list_test.ListInterface = list_test.GetList("single")
	switch t := w3.(type) {
	case Walker:
		fmt.Println("是一个Walker接口")
	case Flyer:
		fmt.Println("是一个Flyer接口")
	//case *Pig:
	//	fmt.Println("是猪")
	//case *Bird:
	//	fmt.Println("是鸟")
	case nil:
		fmt.Println("类型为空")
	default:
		fmt.Printf("未判断类型%T\n", t)
	}

}
