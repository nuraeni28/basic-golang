package main

import "fmt"

func main(){

	//first way of looping
	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	//second way of looping
	j:= 0

	for j < 3{
		fmt.Println(j)
		j++
	}

	//third way of looping
	k:=0
	for{
		fmt.Println(k)

		k++

		if k == 3{
			break //looping berhenti jika sudah memenuhi kondisi
		}
	} 

	//break and continue keywords
	for l:=1;l<10;l++{
		if(l%2 == 1){
			continue //jika angkanya ganjil maka akan diskip prosesnya
		}
		if l > 8{
			break
		}

		fmt.Println("Angka genap : " ,l)
	}

	//nested looping
	for m:=1;m<5;m++{
		for n:=m;n<5;n++{
			fmt.Print(n," ")
		}
		fmt.Println()
	}

	//looping with label

	outerLoop:
		for o := 0; o < 5; o++{
			fmt.Println("Perulangan ke - ", o+1)
			for p := 0 ; p < 5; p++{
				if o == 2 {
					break outerLoop
				}
				fmt.Print(p, "")
			} 
			fmt.Println()
		}

}