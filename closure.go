package main

import (
	"fmt"
	"strings"
)

func main() {
	//declare closure with variable
	var evenNumbers = func(numbers ...int) []int {
		var result []int
		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	numbers := []int{4, 8, 9, 10, 12, 15}

	fmt.Println("Angka genap : ", evenNumbers(numbers...))

	//closure IIFE
	var isPalindrome = func (str string) bool {
		var temp string = ""

		for i := len(str)-1; i >= 0; i--{
			temp += string(byte(str[i]))

		}
		return temp == str
		
	}("katak")

	fmt.Println(isPalindrome)

	//print closure as a return value
	var studentLists = []string{"nuraeni", "taufik", "fasya", "sahar"}

	find := findStudent(studentLists)

	println(find("sahar"))

	//print oddnumbers
	findOdd := findOddNumbers(numbers,func(number int) bool {
		return number%2 != 0
	} )

	fmt.Println("Total angka ganjil : ", findOdd)

	


}

//closure as a return value
	func findStudent(students []string) func(string) string{
		return func(s string) string {
			var student string
			var position int

			for i, v := range students {
				if strings.ToLower(v) == strings.ToLower(s){
				student = v
				position = i
				}
			}

			if student == ""{
				return fmt.Sprintf("%s doesn't exist!!!", s)
			}
			return fmt.Sprintf("We found %s at position %d", s, position)
		}
	}

//closure callback
func findOddNumbers(numbers []int, callback func (int) bool ) int {
	//callback func (int) bool dapat disingkat menjadi alias menjadi : type isOddNum func(int) bool
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v){
			totalOddNumbers++
		}
	}

	return totalOddNumbers
	
}
	


