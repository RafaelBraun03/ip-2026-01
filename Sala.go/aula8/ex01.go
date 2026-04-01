package main

import f "fmt"

func soma(x, y float64) float64{
	return x + y
}

func main() {
	var x, y float64
	f.Println("Insira dois números:")
	f.Scan(&x, &y)

	f.Println("o valor da soma é:",soma(x, y))
}