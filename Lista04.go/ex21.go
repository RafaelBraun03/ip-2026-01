package main

import f "fmt"

func main() {
	var num [10]int
	var invert [10]int
	var n int
	for i:= 0; i <len(num); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&num[i])
	}
	//inverter:
	for i:=0;i<len(num);i++{
		invert[i] = num[9-i]
	}
	f.Println("Insira o código numérico(1 ou 2)")
	f.Scan(&n)
	if n==1{
		f.Println(num)
	}else{
		f.Println(invert)
	}
}