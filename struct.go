package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

//embedded struct
type Person struct{
	name string
	age int
}

type Student struct {
	nim int
	person Person
}

func main() {
	var employee Employee

	employee.name = "Nuraeni"
	employee.age = 24
	employee.division = "IT"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	employee2 := Employee{name: "Taufik", age: 23, division: "Mecanic"}
	fmt.Printf("Employee 2 : %+v\n", employee2)

	//pointer to a struct
	 var employee3 *Employee = &employee
	 fmt.Println("Employee 1 : ", employee.name) 
	 fmt.Println("Employee 3 : ", employee3.name) 

	 employee3.name = "Neni"
	 fmt.Println("Employee 1 : ", employee.name) 
	 fmt.Println("Employee 3 : ", employee3.name) 


	 //embedded struct
	 var student Student
	 student.nim = 42519044
	 student.person.name = "Fasya"
	 student.person.age = 24

	 fmt.Printf("Data mahasiswa : %+v\n", student)

	 //anonymous struct tanpa pengisian field
	 var student1 = struct {
		person Person
		division string
	 }{}
	 student1.person.name = "Sahar"
	 student1.person.age = 23
	 student1.division = "Marketing"

	 //anonymous struct dengan pengisian field
	 var student2 = struct {
		person Person
		division string
	 }{
		person : Person{name: "Nas", age: 10},
		division: "Human Resource",
	 }

	 fmt.Printf("Data mahasiswa 1 : %+v\n", student1)
	 fmt.Printf("Data mahasiswa 2 : %+v\n", student2)

	 //slice of struct
	 var people = []Person{
		{name: "Zalza", age: 23},
		{name: "Kiya", age: 23},
		{name: "Mita", age: 24},
	 }
	 
	 for i, v := range people {
		fmt.Printf("Mahasiswa ke-%d : %+v\n", i+1, v)
	 }

	 //slice of anonymous struct
	 var animals = []struct {
		name string
		tipe string
	 }{
		{name: "chicken", tipe: "omnivore"},{ name:"tiger", tipe:"carnivore"},
		}

	 for i, v := range animals {
		fmt.Printf("Hewan ke-%d : %+v\n", i+1,v )
	 }



}