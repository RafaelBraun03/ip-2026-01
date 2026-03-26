package main

import (f "fmt"
        m "math")

func main() {

	var x string
	var a, v, r, h float64

	f.Println("Selecione a opção que deseja calcular a área e o volume:")
	f.Println("1- cone")
	f.Println("2- cilindro")
	f.Println("3- esfera")
	f.Scan(&x)
	

	if x == "1" {
		f.Println("Digite o raio e a altura: ")
	    f.Scan(&r, &h)
		a = m.Pi*r*m.Sqrt(r*r + h*h)
		v = (m.Pi*r*r*r*h)/3
		f.Printf("O volume e a área são : %.2f , %.2f", v, a)
	}
	if x== "2" {
		f.Println("Digite o raio e a altura: ")
	    f.Scan(&r, &h)
		a= 2*m.Pi*r*h
		v = m.Pi*r*r*h
		f.Printf("O volume e a área são : %.2f , %.2f", v, a)
	}
	if x == "3" {
		f.Println("Digite o raio: ")
	    f.Scan(&r)
		v = (4*m.Pi*r*r*r)/3
		a = (4*m.Pi*r*r)
		f.Printf("O volume e a área são : %.2f , %.2f", v, a)
	}


}