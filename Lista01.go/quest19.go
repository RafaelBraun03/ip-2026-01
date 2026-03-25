package main

import f "fmt"

func main () {

    var k, n int
	var s float32

	f.Println("Digite um número inteiro: ")
    f.Scan(&n)
     s=0
	if n > 1 {
	
		for k = 1; k <= n; k++{
			s = s + (1/float32(k))
		}
	}else {
		f.Println("Número inválido")
	}
	f.Printf("%.6f", s) 
}
