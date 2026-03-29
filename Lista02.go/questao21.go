package main

import f "fmt"

func main () {

	var n, n1, n2, n3, nota, mediaex float32
	var conceito string


    f.Println("Digite o número do aluno: ")
    f.Scan(&n)
	f.Println("Digite o valor primeira nota: ")
    f.Scan(&n1)
	f.Println("Digite o valor segunda nota: ")
    f.Scan(&n2)
	f.Println("Digite o valor terceira nota: ")
    f.Scan(&n3)
	f.Println("Digite o valor da média dos exercícos: ")
    f.Scan(&mediaex)
	nota = (n1 + 2*n2 + 3*n3 + mediaex)/7

    f.Println("Número do aluno:", n)
	f.Println("Nota 1:", n1)
	f.Println("Nota 2:", n2)
	f.Println("Nota 3:", n3)
    if (nota < 4.0) {
		conceito = "E"
	}else {
	   if (nota >= 4.0) && (nota < 6.0) {
		   conceito = "D"
	   }else {
		    if (nota >= 6.0) && (nota < 7.5) {
			   conceito = "C"
		    }else {
			   if (nota >= 7.5) && (nota < 9.0) {
				   conceito = "B"
			   }else {
				   conceito= "A"
			   }
		    }
	   }
	}
	f.Println("NOTA = ", nota, " ", "CONCEITO = ", conceito)
	
	if conceito == "A" || conceito == "B" || conceito == "C" {
		f.Println("STATUS: APROVADO")
	}else {
		f.Println("STATUS: REPROVADO")
	}
}