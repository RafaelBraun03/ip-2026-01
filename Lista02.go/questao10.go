package main

import f "fmt"

func main () {

	var destino, volta string

	f.Println("Qual o destino da viagem?")
	f.Scan(&destino)

	f.Println("A viagem é de ida e volta?")
	f.Scan(&volta)


	if destino == "1" && volta == "sim" {
	   f.Println("O preço da passagem é 900.00")
	}
	if destino == "1" && volta == "nao" {
	   f.Println("O preço da passagem é 500.00")
	}
	if destino == "2" && volta == "sim" {
	   f.Println("O preço da passagem é 650.00")
	}
	if destino == "2" && volta == "nao" {
	   f.Println("O preço da passagem é 350.00")
	}
	if destino == "3" && volta == "sim" {
	   f.Println("O preço da passagem é 600.00")
	}
	if destino == "3" && volta == "nao" {
	   f.Println("O preço da passagem é 350.00")
	}
	if destino == "4" && volta == "sim" {
	   f.Println("O preço da passagem é 550.00")
	}
	if destino == "4" && volta == "nao" {
	   f.Println("O preço da passagem é 300.00")
	}
}