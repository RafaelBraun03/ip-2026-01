package main

import (
	f "fmt"
	"math"
)

func main(){
		var num [100]int

	for i:= 0; i <len(num); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&num[i])
	}
	s:= 0.0
	for j:=0;j<50;j++{
		s = s + (math.Pow(float64(num[j]-num[99-j]),3))
	}
	f.Println("S=",s)
}