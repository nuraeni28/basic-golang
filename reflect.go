package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string
	age int
}
func (s *student) SetName(name string){
		s.name = name
	}

func main() {
	var number = 28
	var reflectValue = reflect.ValueOf(number)

	//check type data with method kind
	if reflectValue.Kind() == reflect.Int{
		fmt.Println("nilai variabel : ", reflectValue.Int())
	}
	fmt.Println("tipe data number : " , reflectValue.Type())


	//accessing value using interface method
	fmt.Println("nilai variabel with interface method : ", reflectValue.Interface())

	//identifiying method information
	var s1 = &student{name : "neni", age : 24 }
	var s1ReflectValue = reflect.ValueOf(s1)
	method := s1ReflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("taufik"),
	})
	fmt.Println("nama", s1.name)
	fmt.Println("s1 reflect value", s1ReflectValue)
	fmt.Println("name value after call with method", s1.name)
	
}