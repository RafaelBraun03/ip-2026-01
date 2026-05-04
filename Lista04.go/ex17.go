package main

import f "fmt"

func primo(n int)bool {

	for i:=2;i*i<=n; i++{
		if n % i==0 {
			return  false
		}
	}
	return true
}
func main(){
	var num [10]int

	for i:= 0; i <len(num); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&num[i])
	}
	f.Println("-------------------------")
	f.Println("São primos:")
	for i,v:= range num{
		if primo(v)==true{
			f.Printf("Número: %d     Índice:%d", v,i)
			f.Println("")
		}
	}
}