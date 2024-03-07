package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("main execution start")
	//tambahkan waitgroup
	var wg sync.WaitGroup
	wg.Add(1)
	go firstProcessWait(3, &wg)

	secondProcessWait(3)
	
	fmt.Println("no of goroutines:", runtime.NumGoroutine()) 
	wg.Wait() //untuk menunggu seluruh go routine menyelesaikan prosesnya

	fmt.Println("main execution end")

}

func firstProcessWait(index int, wg *sync.WaitGroup){
	fmt.Println("firstProcess start")

	for i:= 0; i < index; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("firstProcess end")

	wg.Done()  //harus ditambahkan done

}
func secondProcessWait(index int){
	fmt.Println("secondProcess start")

	for j:= 0; j < index; j++ {
		fmt.Println("j = ", j)
	}

	fmt.Println("secondProcess end")

}