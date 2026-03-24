package main

import "fmt"

func main () {

	var n1, n2, p1, p2, media float64

	fmt.Println("Digite a primeira nota: ")
	fmt.Scan(&n1)

	fmt.Println("Digite a segunda nota: ")
	fmt.Scan(&n2)

	fmt.Println("Digite os pesos da primeira e da segunda nota: ")
	fmt.Scan(&p1, &p2)

	media = (n1*p1 + n2*p2)/(p1 + p2)
	fmt.Println(media)
}