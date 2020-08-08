package main

import (
	"fmt"
	"pruebatecnica/features/purchase/domain/models"
)

func main() {
	p := Buyer{
		id:   "2222",
		name: "Alice",
		age:  "22",
	}
	fmt.Println(p)

}
