package main

import f "fmt"

func main() {
	var n int

	f.Println("Insira um número inteiro:")
	f.Scan(&n)
	triangular := false
	for i := 1; i*(i+1)*(i+2)<=n; i++{
	  if n > 0{	
		if i*(i+1)*(i+2) == n{
			triangular = true
			break
		}
	  }else {
		f.Println("número inserido inválido")
	  }
	}
		if triangular == true {
			f.Println("O número é triangular")
		}else{
			f.Println("O número não é triangular")
		}
}