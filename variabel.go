package main

import "fmt"

func main() {
	//variable with data type
	// var name string = "Nuraeni"
	// var age int = 24

	// fmt.Println("Nama : ", name)
	// fmt.Println("Usia : ", age)

	//variable without data type
	// var nickName = "neni"
	// var height = 155

	// fmt.Printf("%T, %T", nickName, height) //untuk cek type data

	//short declaration
	// favoriteFood := "noodles" //teknik ini hanya bisa digunakan jika kita langsung memberikan nilai pada variabel tersebut

	// fmt.Printf("%T", favoriteFood);

	//multiple variable declaration
	var student1, student2, student3 string = "neni", "taufik", "fasya"

	var age1, age2, age3 int

	age1, age2, age3 = 24, 23, 23

	println(student1, student2, student3)
	println(age1, age2, age3)

	//multiple variable with different type
	var fullName, age, address = "Nuraeni", 24, "Jalan Hartaco Indah"

	first, second, third := 1, "2", 3

	var test string //pada golang tidak ada variabel yang tidak digunakan
	_ = test        //untuk mengatasi variabel yang belum atau tidak digunakan yaitu dengan menambahkan tanda _ (underscore)

	println(fullName, age, address)
	println(first, second, third)

	//penggunaan printf

	var nama = "Nuraeni"
	var usia = 24
	var alamat = "Jl Hartaco Indah"

	fmt.Printf("Halo nama saya %s, umurku %d dan aku tinggal di %s \n", nama, usia, alamat)

	//multi line

	introduction := `Halo, nama saya Nuraeni
	Kuliah di Politeknik Negeri Ujung Pandang` 

	fmt.Println(introduction)

	angka1 := 2
	angka2 := 5

	sum := float64(angka1 + angka2)
    avg := sum / 2
	fmt.Printf("%.1f\n", avg)

}