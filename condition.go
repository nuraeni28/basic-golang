package main

import "fmt"


func main(){
	//if-else
	var currentYear = 2024

	if age:= currentYear-1999; age < 17 {
		fmt.Println("kamu belum cukup umur untuk membuat ktp")
	} else{
		fmt.Println("kamu sudah bisa membuat ktp")
	}

	//switch case
	var ipk = 2.97

	switch{
	case ipk > 3:
		fmt.Println("perfect")
	default:
		fmt.Println("good")
	}

	//switch fallthrough keyword
	var score = 6

	switch{
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not band")
		fallthrough //fungsi fallthrough berfungsi untuk tetap melanjutkan pengecekan ke case selanjutnya walaupun sudah memenuhi di case ini
	case score > 5:
		fmt.Println("it's ok, but please study harder")
	default:{
		fmt.Println("study harder")
		fmt.Println("you don't have a good score yet")
	}
	}

	//nested conditions
	var price = 90000

	if price >= 100000{
		switch price{
		case 100000:
			fmt.Println("standart")
		default:
			fmt.Println("expensive")
		
	}
}else {
		fmt.Println("cheap")
	}
	
	

}