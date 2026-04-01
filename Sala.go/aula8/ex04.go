package main

import f "fmt"

func fatorial(n int)int {
     var P int = 1
      
	for i := 1; i <= n; i++{
          P = P*i
	}
	return P
}
func main() {
	var n int
    
	f.Println("Digite o número para calcular seu fatorial:")
	f.Scan(&n)
	f.Println("O fatorial do número é:", fatorial(n))
}