package main

import f "fmt"

func main () {

	var r, h , custo float32

   f.Println("Digite o raio da lata: ")
   f.Scan(&r)
   
   f.Println("Digite a altura da lata: ")
   f.Scan(&h)

   custo = 100.00 * (2*3.14159*r*r + 2*3.14159* r * h)

    f.Printf("O valor do custo é = %.2f", custo)
}