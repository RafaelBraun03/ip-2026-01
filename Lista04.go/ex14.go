package main

import f "fmt"

func main(){
	var vetor1 [10]int
	var vetor2 [10]int
	var resultante [20]int
	
	for i:= 0; i <len(vetor1); i++ {
		f.Printf("Digite o %dº valor do vetor 1:", i+1)
		f.Scan(&vetor1[i])
	}
	f.Println("")
	for i:= 0; i <len(vetor2); i++ {
		f.Printf("Digite o %dº valor do vetor 2:", i+1)
		f.Scan(&vetor2[i])
	}
	for i:=0; i<10;i++{
		resultante[2*i] = vetor1[i]
		resultante[(2*i)+1] = vetor2[i]
	}
	f.Println(resultante)

}