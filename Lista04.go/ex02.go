package main

import f "fmt"

func main () {
		var num10 [10]int
		var num5 [5]int

		soma := 0
		f.Println("Primeiro vetor:")
	for i:= 0; i <len(num10); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&num10[i])
	}
		f.Println("Segundo vetor:")
	for i:= 0; i <len(num5); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&num5[i])
		soma = soma + num5[i]
	}
	var primeiro []int 
	var segundo []int
		for _,v := range num10 {
		if v % 2 ==0{
			primeiro = append(primeiro, v+soma)
		}else{
			segundo = append(segundo, v+soma)
		}
	}
	f.Println("Primeiro vetor resultante:", primeiro)
	f.Println("Segundo vetor resultante:", segundo)
}