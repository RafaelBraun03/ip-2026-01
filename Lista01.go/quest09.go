package main

import f "fmt"

func main () {

	var a, b ,c , x float32

	f.Println("Digite o coeficiente A: ")
    f.Scan(&a)
	f.Println("Digite o coeficiente B: ")
    f.Scan(&b)
	f.Println("Digite o coeficiente C: ")
    f.Scan(&c)
	 x = b*b - 4*(a*c)

	 f.Printf("O valor do delta é = %.2f ", x )
}