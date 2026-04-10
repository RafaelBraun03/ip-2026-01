package main

import "fmt"

type Criatura struct {
	Nome string
}

func main() {
	c := Criatura{
		Nome: "Jac and Haniel",
	}
	fmt.Println(c.Nome)
}
