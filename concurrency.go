package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("main execution start")
	go firstProcess(3)

	secondProcess(3)


	fmt.Println("no of goroutines:", runtime.NumGoroutine()) 
	//ketika dijalankan maka banyakan goroutinesnya ada 2 yaitu 1 dari main default dan satunya lagi adalah goroutines firstProcess
	
	time.Sleep(time.Second * 2)
	fmt.Println("main execution end")

}

func firstProcess(index int){
	fmt.Println("firstProcess start")

	for i:= 0; i < index; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("firstProcess end")

}
func secondProcess(index int){
	fmt.Println("secondProcess start")

	for j:= 0; j < index; j++ {
		fmt.Println("j = ", j)
	}

	fmt.Println("secondProcess end")

}