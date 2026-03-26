package main

import f "fmt"

func main () {

	var n1, n2, n3, media float32

	 f.Print("Digite as 3 notas: ")
	 f.Scan(&n1, &n2, &n3)

	 media = (n1 + n2 + n3)/3
	 f.Println("A média é:", media)

	 if media >= 7 {
		f.Println("Status: APROVADO")
	 }else {
		f.Println("Status: REPROVADO")
	 }

}