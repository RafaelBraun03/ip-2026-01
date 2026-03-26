package main

import ( "fmt"
         "math")

func main () {

	var n, k, s float64

	fmt.Println("Insira um número:")
	fmt.Scan(&n)

	k = math.Sqrt(n)
	s = math.Pow(n, 2)

	if n >= 0 {
		fmt.Println(k)
	}else {
		fmt.Println(s)
	}

}		