package main

import f "fmt"

func main () {

	var n int

	 f.Println("Digite um número inteiro: ")
     f.Scan(&n)
 
     if (n % 3 == 0) && (n % 5 == 0) {
     f.Println("O NÚMERO É DIVISÍVEL")
	 }else{
     f.Println("O NÚMERO NÃO É DIVISÍVEL")
	 }
}