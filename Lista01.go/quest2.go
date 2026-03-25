package main

import f "fmt"

func main () {

	var per_popular, per_geral, per_arquibancada, per_cadeiras, renda float64
	var ingressos int

	f.Println("Digite o número de ingressos: ")
    f.Scan(&ingressos)

    f.Println("Digite a porcentagem de pessoas na categoria popular: ")
    f.Scan(&per_popular)

    f.Println("Digite a porcentagem de pessoas na categoria geral: ")
    f.Scan(&per_geral)

    f.Println("Digite a porcentagem de pessoas na categoria arquibancada: ")
    f.Scan(&per_arquibancada)

    f.Println("Digite a porcentagem de pessoas na categoria cadeiras: ")
    f.Scan(&per_cadeiras)

	renda = (per_popular*1.0 + per_geral*5.0 + per_arquibancada*10.0 + per_cadeiras*20.0)/100.0


	f.Println(renda * float64(ingressos))

}	