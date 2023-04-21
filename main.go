package main

import (
	"fmt"
)

type Author struct {
	Name string
	age  int
	sal  int
	city string
}
type Book struct {
	Name  string
	pages int
	cost  int
}

func main() {
	a1 := new(Author)

	a1.Name = "Chand"
	a1.age = 50
	a1.sal = 400000
	a1.city = "Hyd"

	fmt.Println(a1)

	a2 := Author{"Krishna", 30, 300000, "Vizag"}
	fmt.Println(a2)

	var a3 = Author{
		Name: "Vijay",
		age:  40,
		sal:  100000,
		city: "Mumbai",
	}
	fmt.Println(a3)

	b1 := new(Book)

	b1.Name = "Magic book"
	b1.pages = 50
	b1.cost = 400

	fmt.Println(b1)

	b2 := Book{"Golang", 100, 500}
	fmt.Println(b2)

	var b3 = Book{
		Name:  "Computers",
		pages: 200,
		cost:  1000,
	}
	fmt.Println(b3)

}
