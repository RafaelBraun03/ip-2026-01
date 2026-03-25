package main

import f "fmt"

func main () {

	var a, r, n, S, i float64

	f.Println("Digite três números: "," ", " ")
    f.Scan(&a, &r, &n)

	for i = 1; i <= n; i++ {
		S = S + a + (i - 1)*r
	}
	f.Println(S)
}