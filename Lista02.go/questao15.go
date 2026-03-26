package main

import f "fmt"

func main () {

	var d, m, a int
	

	f.Println("Digite o dia: ")
	f.Scan(&d)
	f.Println("Digite o mês: ")
	f.Scan(&m)
	f.Println("Digite o ano: ")
	f.Scan(&a)

	mes := []string{
          " ", "Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
	}

	f.Printf("%d de %s de %d", d, mes[m], a)


}