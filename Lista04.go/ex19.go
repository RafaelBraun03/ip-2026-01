package main

import f "fmt"

func main () {
		var num [10]int
		var divisores [5]int

		f.Println("Primeiro vetor:")
	for i:= 0; i <len(num); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&num[i])
	}
		f.Println("Segundo vetor:")
	for i:= 0; i <len(divisores); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&divisores[i])
	}
	for _,v:= range num{
			f.Printf("Número %d",v)
				f.Println()
		for j,x:= range divisores{
			if v % x ==0{
				f.Printf("     Divisível por %d na posição %d", x,j)
				f.Println()
			}else{
				f.Println("Não apresenta divisores no vetor")
			}
		}
	}
}