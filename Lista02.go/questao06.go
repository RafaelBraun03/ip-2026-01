package main

import f "fmt"

func main () {

	var a, b int

	f.Println("Insira dois números inteiros: A e B")
	f.Scan(&a, &b)

	if a%b == 0 {
		f.Println("A é divisível por B")
	}else {
		f.Println("A não é divisível B")
	}
}