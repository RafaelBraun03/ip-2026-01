package main

import f "fmt"

func main() {
	var soma, cont int
	var media float64
    soma = 0
	cont = 0
	i:=50
	for n:=1; n < 20;n++{
		 if (n+50) % 2==0 {
			   cont = cont + 1
			   soma = soma + n + i
	    }
	 }
	 f.Println("A soma é : ",soma)
	 media = float64(soma)/float64(cont)
	 f.Println("A média é:", media)
}
