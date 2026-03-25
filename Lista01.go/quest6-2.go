package main

import f "fmt"

func main () {

var n, i int	
var F, c float64 

f.Println("Digite quantas temperaturas serão convertidas:")
f.Scan(&n)

if n <= 0 {
	f.Println("Quantidade inválida")
}

for i = 1; i <= n; i++ {
	f.Println("Digite a temperatura ", i, ":")
	f.Scan(&F)
	c = (F - 32)*5/9
	f.Printf("A temperatura %d em celcius é: %.2f ", i, c)
	f.Println(" ")
}
}