package main

import f "fmt"
        
func main() {
	var x, y, r float64

	f.Println("Insira a base e o expoente:")
	f.Scan(&x, &y)

    r = 1

	for i:=0; i<int(y); i++{
		r = r*x
	}
	f.Printf("%.4f^%.4f=%.4f", x, y, r)
		
}