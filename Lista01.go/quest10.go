package main

import f "fmt"

func main () {

	var a, b, c, d , Det float32

	f.Println("Digite o elemento a da matriz: ")
    f.Scan(&a)
	f.Println("Digite o elemento b da matriz: ")
    f.Scan(&b)
	f.Println("Digite o elemento c da matriz: ")
    f.Scan(&c)
	f.Println("Digite o elemento d da matriz: ")
    f.Scan(&d)

	Det = a*d - b*c

	 f.Printf("O VALOR DO DETERMINANTE É = %.2f", Det)
}