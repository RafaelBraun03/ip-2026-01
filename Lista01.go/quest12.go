package main

import f "fmt"

func main () {

	var valor, horas, x, y int
	
	

	f.Println("Digite o número de horas que o locatário utilizou a charrete: ")
     f.Scan(&horas)
      
	 x = horas/3
	 y = horas % 3

	  valor = x*10  + y*5.00
 
     f.Printf("O VALOR A PAGAR É = %.2d", valor)
 
}