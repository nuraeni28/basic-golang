package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("nuraeni", 24)
	var names = []string{"Muh ", "Taufik Witri"}
	var msg = bio(names, 23, "ternate")
	
	fmt.Println(msg)

	var diamater float64 = 15
	var area, circumference = calculate(diamater) //cara pemanggilan returning multiple values
	fmt.Println("Luas : ", area)
	fmt.Println("Keliling : ", circumference)

	a := 4.5
	t:= 2.0

	var luasSegitiga = triangleArea(a,t)

	fmt.Println("Luas segitiga : ", luasSegitiga)

	//variadic print
	studentList := print("Aril", "Jasmine", "Frozen")

	fmt.Printf("%v\n", studentList)

	//variadic 2
	numberList := []int{1,2,3,4,5,6,7,8,9,10}

	total := sum(numberList...)

	fmt.Println("Total : ", total)

	//variadic 3
	profile("Nuraeni", "Noodles", "Fried Rice", "Coto")

}

func greet(name string, age int8) {
	fmt.Println("Hello guys, my name is : ", name)
	fmt.Println("my age is : ", age)
}

//function return
func bio(names []string, age int8, address string) string{ // string setelah () merupakan type data return nya
	var fullName = strings.Join(names, "")
	
	var result string = fmt.Sprintf("Hello, let me introduce my self. My name is %s, my age is %d and i live at %s", fullName, age, address)

	return result
}

//returning multiple values
func calculate(d float64) (float64, float64){
	//hitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	//hitung keliling
	var circumference = math.Pi * d 

	return area, circumference

}

//predifined return value
func triangleArea(a float64, t float64) (area float64){
	area = (a * t)/2

	return area
}

//variadic function (function yang dapat menerima argumen yang tak terbatas jumlahnya)
func print(names ...string)[]map[string]string{ //variadic funct ditandai dengan ...string
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

//variadic function 2
func sum(numbers ...int) int{
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total

}

//variadic function 3
func profile(name string, favoriteFoods ...string){ //parameter biasa bisa digabungkan dengan parameter variadic. Namun, paramater variadic harus diletakkan di akhir parameter
	mergeFavFoods := strings.Join(favoriteFoods, ",")

	fmt.Println("Hello guys, my name is : ", name)
	fmt.Println("My favorite foods are", mergeFavFoods)

}