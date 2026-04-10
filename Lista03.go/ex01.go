package main

import (f "fmt"
        m "math")

func main() {
	var x, y, r float64

	f.Println("Insira a base e o expoente:")
	f.Scan(&x, &y)

    r = m.Pow(x, y)
	f.Printf("%.4f^%.4f=%.4f", x, y, r)
}		