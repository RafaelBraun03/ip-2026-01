package main

import f "fmt"

func main () {

	var id int

	f.Println("Digite a idade")
	f.Scan(&id)

	if id < 0 {
		f.Println("Idade inválida")
	}
	if id >= 0 && id <= 2 {
		f.Println("Recém-nascido")
	}
	if id >= 3 && id <=11 {
		f.Println("Criança")
	}
	if id >= 12 && id <= 19 {
		f.Println("Adolescente")
	}
	if id >= 20 && id <= 55 {
		f.Println("Adulto")
	}
	if id > 55 {
		f.Println("Idoso")
	}
}