package main

import f "fmt"

func main () {

	var n1, n2, s int

	f.Println("Insira dois números inteiros: ")
	f.Scan(&n1, &n2)

	s = n1 + n2

	if s > 20 {
		f.Println(s + 8)	
	}else {
		f.Println(s - 5)
	}




}