package main

import f "fmt"

func main () {

	var x,y,n int

	f.Println("Digite dois números: ")
	f.Scan(&x, &y)

	if x % 2 == 0 {
		for n = 1; n <= y; n++ {
			f.Print(x, " ")
			x = x + 2
		}
	}else {
		f.Println("O primeiro número é impar")
	}
}