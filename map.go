package main

import "fmt"

func main(){
	var person map[string]string //deklarasi
	//map[string]string berarti tipe data  key dari map harusstring dan value dari map juga harus string.
	
	person = map[string]string{} //inisialisasi

	person["name"] = "nuraeni"
	person["age"] = "24"
	person["address"] = "Hartaco Indah"

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])

	//directly assigns values and keys to the map
	var student = map[string]string{
		"name" : "nuraeni",
		"gender" : "female",
		"nim" : "42519044",
	}
	fmt.Println("name: ", student["name"])
	fmt.Println("gender: ", student["gender"])
	fmt.Println("nim: ", student["nim"])

	//looping with map
	for k, v := range student {
		fmt.Println(k, ":", v)
	}

	//deleting value
	fmt.Println("Before delete : ", student)

	//detecting value
	value, exist := student["gender"]  //Variable value yang kita berikan akan mengembalikan value asli darimap nya jika memang key nya ada, jika tidak ada maka kita akanmendapat zero value dari value nya
	if exist {
		fmt.Println(value)
	}else{
		fmt.Println("value doesn't exist")
	}

	delete(student, "gender") //hapus map
	value, exist = student["gender"]  //Variable value yang kita berikan akan mengembalikan value asli darimap nya jika memang key nya ada, jika tidak ada maka kita akanmendapat zero value dari value nya
	if exist {
		fmt.Println(value)
	}else{
		fmt.Println("value has been deleted")
	}

	fmt.Println("After delete", student)

	//combining slice with map
	students := []map[string]string{
		{"name" : "taufik witri", "gender" : "male", "nim" : "42519041"},
		{"name" : "nuraeni", "gender" : "female", "nim" : "42519044"},
		{"name" : "fasya", "gender" : "male", "nim" : "42519038"},

	}

	for i, s := range students {
		fmt.Printf("Index: %d, name: %s, gender: %s, nim: %s\n ", i, s["name"], s["gender"], s["nim"])
	}

	


}