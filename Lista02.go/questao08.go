package main

import f "fmt"

func main () {

	var n int

	f.Println("Insira um número inteiro: ")
	f.Scan(&n)

	if n > 20 && n < 90 {
		f.Println("Está entre 20 e 90")
	}else {
		f.Println("Não está entre 20 e 90")
	}
}