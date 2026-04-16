package main

import f "fmt"

func expoente(x, n int)int{
	if n == 0 {
		return 1
	}
	return x*expoente(x, n-1)
}

func main() {
	var x, n int
	f.Println("Insira a base e o expoente:")
	f.Scan(&x, &n)

	f.Println(expoente(x,n))
}