package main

import f "fmt"

func main () {

	var n, x int

	f.Println("Digite o número N: ")
	f.Scan(&n)

	for x = 2; x <= n; x++ {
		if (x % 2 == 0) && (n > 5) && (n < 2000) {
			f.Println(x, "^2 = ", x*x)
		}
	}
}