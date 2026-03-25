package main

import f "fmt"

func main () {

	var salario, final float64

     f.Println("Digite o salário do funcionário: ")
     f.Scan(&salario)

	 if salario <= 300.00 {
		final = salario*1.5
	 }
	 if salario > 300.00 {
		final = salario*1.3
	 }
	 f.Printf("O salário com reajuste é: %.2f",final)
}