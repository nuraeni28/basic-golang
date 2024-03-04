package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}

func (r rectangle) area() float64{
	return r.height * r.width
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64{
	return 2 * (r.height + r.width)
}






func main() {
	var c1 shape = circle{radius: 7}
	var r1 shape = rectangle{width: 10, height: 5}

	//empty interface
	var randomValues interface{}

	_ = randomValues

	randomValues = 123

	randomValues = "Nuraeni"

	randomValues = true

	randomValues = []string{"neni","taufik"}

	//empty interface (type assertion)
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	//empty interface (map, slice with empty interface)
	rs := []interface{}{1,"neni", "hartaco indah", true}

	rm := map[string]interface{}{
		"nama" : "nuraeni",
		"status" : true,
		"age" : 23,
	}

	fmt.Printf("Tipe c1 : %T\n", c1)
	fmt.Println("Luas lingkaran : ", c1.area())
	fmt.Println("Keliling lingkaran : ", c1.perimeter())
	fmt.Printf("Tipe r1 : %T\n", r1)
	fmt.Println("Luas persegi panjang : ", r1.area())
	fmt.Println("Keliling persegi panjang : ", r1.perimeter())

	fmt.Println("Hasil perkalian : ", v)
	fmt.Println("Hasil slice : ", rs)
	fmt.Println("Hasil map : ", rm)

}


