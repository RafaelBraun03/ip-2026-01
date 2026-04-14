package main

import f "fmt"

func main() {
	var n , a, b int
	f.Println("Digite os dois primeiros termos:")
	f.Scan(&a, &b)// faremos para a=1 e b=2
	f.Println("Insira a quantidade de termos da série de Fetuccine que serão impressa:")
	f.Scan(&n)
	if n < 3 {
		f.Println("Número inválido")
		return
	}
	f.Print("1 2 ")
	for i := 1; i <= n; i++{
	var	proximo int 
		if i%2==0{
			proximo = b-a
		}else {
			proximo=b+a
		}
	  f.Print(proximo, " ")
	  a, b = b, proximo
	}
	f.Println()
}