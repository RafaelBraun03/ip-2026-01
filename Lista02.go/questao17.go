package main

import f "fmt"

func main () {

	var u  int
    var categoria string
    var consumo, valor float32

	f.Println("Digite o identificador da conta: ")
    f.Scan(&u)
 
    f.Println("Digite o consumo de água em metros cúbicos: ")
    f.Scan(&consumo)
 
    f.Println("Digite o tipo de consumidor: ")
    f.Scan(&categoria)

	if categoria == "R" {
		valor = 5.0 + 0.05*consumo
	}else {
		if categoria == "C" {
			if consumo <= 80.0 {
				valor = 500.0
			}else {
				valor = 500 + 0.25*(consumo - 80)
			} 		
		}else {
			if categoria == "I" {
				if consumo <= 100.0 {
					valor = 800.0
				}else {
					valor = (consumo - 100.0)*0.04
				}
			}
		}
	}
	f.Println("CONTA = ", u)
	f.Println("VALOR DA CONTA = ", valor)
}