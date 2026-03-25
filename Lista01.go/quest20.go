package main

import f "fmt"

func main () {

	var  h, m, s, T int

	 f.Println("Digite o valor em horas: ")
     f.Scan(&h)
	 f.Println("Digite o valor em minutos: ")
     f.Scan(&m)
	 f.Println("Digite o valor em segundos: ")
     f.Scan(&s)

	 T = h*3600 + m*60 + s
	 f.Println("O TEMPO EM SEGUNDOS É: ", T)
}
