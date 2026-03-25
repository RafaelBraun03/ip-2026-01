package main

import f "fmt"

func main () {

	var nota float32
	var conceito string

	f.Println("Digite o valor da nota: ")
    f.Scan(&nota)

	if (nota >= 0.0) && (nota < 6.0) {
		conceito = "D"
	}else {
		 if (nota >= 6.0) && (nota < 7.5) {
			conceito = "C"
		 }else {
			if (nota >= 7.5) && (nota < 9.0) {
				conceito = "B"
			}else {
				conceito= "A"
			}
		 }
	}
	f.Println("NOTA = ", nota, " ", "CONCEITO = ", conceito)
}