package main

import (
	"fmt"
)

func main(){

	var numbers [4]int
	numbers = [4]int{1,2,3,4}

	var strings = [3]string{"Neni", "Nina", "Nas"}

	fmt.Printf("%#v\n",numbers) // %#v untuk melihat panjang dan isi arraynya
	fmt.Printf("%#v\n",strings)


	//modify element through index
	strings[0] = "Nuraeni"
	fmt.Printf("Array setelah modifikasi : %#v\n",strings)

	//loop through elements
	animals:= [3]string{"duck", "chicken", "bird"}
	fmt.Printf("%#v\n",animals)

	//first way
	for i, v := range animals{
		fmt.Printf("Index: %d, Value: %s\n", i,v)
	}

	//second way
	for i:=0; i < len(animals); i++{
		fmt.Printf("Index: %d, Value: %s\n", i, animals[i])
	}


	//multidimensional array
	balances:= [2][3]int{{1,2,3},{4,5,6}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d", value)
		}
		fmt.Println()

	}

	






}