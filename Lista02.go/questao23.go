package main

import f "fmt"

func main () {

	var id int

	f.Println("Digite a idade")
	f.Scan(&id)

	if id < 0 {
		f.Println("Idade inválida")
	}
	if id >= 0 && id < 16 {
		f.Println("Não-eleitor")
	}
	if id > 18 && id <65 {
		f.Println("Eleito obrigatório")
	}
	if id >=16 && id <= 18 || id >= 65 {
		f.Println("Eleitor facultativo")
	}
}