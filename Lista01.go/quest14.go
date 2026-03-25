package main

import  ( "fmt"
         "math")

func main () {

	var h, a, volume float64

	fmt.Println("Digite a altura da pirâmide: ")
    fmt.Scan(&h)
   
    fmt.Println("Digite a aresta da base hexagonal: ")
    fmt.Scan(&a)
	volume = (h * math.Pow(a, 2) * math.Sqrt(3))/2

	fmt.Printf("O VOLUME DA PIRAMIDE: %.2f METROS CÚBICOS", volume)
}