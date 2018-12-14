/*
*
* Go 系列教程 22.信道
*
 */

package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}
func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}
