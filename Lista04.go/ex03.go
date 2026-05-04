package main

import f "fmt"

func main() {
	var num [10]int

	for i:= 0; i <len(num); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&num[i])
	}
	soma:=0
	qtd := 0
	var pares []int
	var impares []int
	for _,v:= range num{
		if v%2==0 {
			pares = append(pares, v)
			soma += v
		}else {
			impares = append(impares, v)
			qtd += 1
		}
	}
	f.Println("Pares:", pares)
	f.Println("Soma dos pares:",soma)
	f.Println("Impares:", impares)
	f.Println("Quantidade impares:", qtd)
}