package main

import f "fmt"

func main () {

	var n1, n2, n3, N, K int

   f.Println("Digite n1: ")
   f.Scan(&n1)
   f.Println("Digite n2: ")
   f.Scan(&n2)
   f.Println("Digite n3: ")
   f.Scan(&n3)
   N = n1*100 + n2*10 + n3 
   K = N*N

   if (n1 <= 0) || (n1 >9) || (n2 < 0) || (n2 > 9) || (n3 < 0) || (n3 > 9) {
	f.Println("Digito inválido")
   }else {
	f.Println(N, ",", K)
   }

}