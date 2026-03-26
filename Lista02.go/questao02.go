package main

import f "fmt"

func main () {

	var n int
     f.Println("Digite um número inteiro:")
	 f.Scan(&n)

	 if n > 0 {
        f.Println("O número é positivo")
	 }
	 if n == 0 {
		f.Println("O número é nulo")
	 }
	 if n < 0 {
		f.Println("O número é negativo")
	 }
}