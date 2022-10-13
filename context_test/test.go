package context_test

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Main() {
	fmt.Println("main start")
	// context类似套娃，最顶层的context要套一个空context进去，下一层的context可以把上一层context套进去，这样后面就可以获取到当前层context内容到最顶层的内容
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10) // 最顶层套一个空context，无值、无过期时间、无取消方法。自己本身有一个10秒的过期时间
	ctx2 := context.WithValue(ctx, "a", "b")                            // 继承了上层context的过期时间10秒
	ctx3, cancelFunc3 := context.WithTimeout(ctx2, time.Second*8)       // 继承了上层context的值 a => b，同时自己也有8秒的过期时间
	go testContext(ctx3, time.Second*6)
	log.Println("main finish")
	var out = false
	for !out {
		select {
		case <-ctx.Done(): //ctx未到期，继续等待
			time.Sleep(time.Second)
			log.Println("main等待到期了")
			cancelFunc3()
			out = true
			break
		case <-time.After(time.Second * 5): // 等待5秒后，取消ctx3的运行
			log.Println("等待了5秒，要求ctx3停止运行")
			cancelFunc3()
		}
	}
}

func testContext(ctx context.Context, sec time.Duration) {
	dl, _ := ctx.Deadline()
	log.Println("test start，到期时间", dl.Sub(time.Now()).Seconds(), "包含内容a:", ctx.Value("a"))
	select {
	case <-ctx.Done(): // 等待ctx结束
		log.Println("test is err", ctx.Err())
	}
}
