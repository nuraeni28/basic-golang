package main

import "fmt"

func main(){
	var fruits = make([]string, 3)

	//append function
	fruits[0] = "apple"

	fmt.Printf("%#v\n", fruits)

	//appends way
	var animals = make([]string,3)
	animals = append(animals, "pig", "dog", "chicken")
	
	
	fmt.Printf("%#v\n", animals)

	//append fuction with ellipsis

	var films1 = []string{"baby boss", "fast furious"}
	var films2 = []string{"agak laen", "spiderman"}

	var appendFilms = append(films1, films2...)

	fmt.Printf("%#v\n", appendFilms)

	//copy function
	copyFilms := copy(films1, films2)

	fmt.Println("Film1 => ", films1)
	fmt.Println("Film2 => ", films2)
	fmt.Println("Copied elements", copyFilms)


	//slicing
	//penggunaan slice lebih hemat memori jika dibandingkan dengan tipe data array

	students := []string{"neni", "taufik", "fasya", "sahar"}

	var students1 = students[1:3] //diambil dari index 1 sampai sebelum 3
	var students2 = students[0:] //mengambil dari index 0 hingga akhir
	var students3 = students[:] //sama dengan students[:len(students)] ini untuk menampilkan semua index

	fmt.Printf("%#v\n", students1)
	fmt.Printf("%#v\n", students2)
	fmt.Printf("%#v\n", students3)

	//combining slicing and append
	students = append(students[:3],"zahrah" )

	fmt.Printf("%#v\n", students)

	//backing array
	jobs := []string{"doctor", "engineer", "police", "data analyst"}

	job1 := jobs[2:4]
	job1[0] = "lecturer" //jika kita mengganti value dari index di job1 maka value di jobs juga akan terganti karena jobs dan job1 masih dalam 1 backing array

	fmt.Println("Jobs => ", jobs)
	fmt.Println("Job 1 => ", job1)

	//cap function
	fmt.Println("students cap : ", cap(students))
	fmt.Println("students len : ", len(students))

	fmt.Println("students1 cap : ", cap(students1))
	fmt.Println("students1 len : ", len(students1))

	fmt.Println("students2 cap : ", cap(students2))
	fmt.Println("students2 len : ", len(students2))

	//creating a new backing array
	cars := []string{"Honda", "Ford", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:3]...)
	newCars[0] = "Nissan" //Ketika index ke-0 dari newCars dirubah, maka cars tidak ikut terubah dikarenakan mereka tidak memiliki backing array yang sama

	fmt.Println("Cars : ", cars)
	fmt.Println("New Cars : ", newCars)

		
}