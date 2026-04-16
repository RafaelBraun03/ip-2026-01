package main

import f "fmt"

func soma(s []float64, n int)float64{
	if n == 0 {
		return s[n]
	}	
		return s[n] + soma(s, n-1)
}
func main() {
	var s = []float64 {1, 2, 3, 4, 5}

	f.Println(soma(s, len(s) - 1))
}