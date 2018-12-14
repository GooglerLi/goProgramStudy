/*
*
* Go 系列教程 22.信道
*
 */

package chanstudy

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

// ChanExample 信道学习例子
func ChanExample() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}
