package main

import (
	f "fmt"
	"math"
)

func main() {
	var num [15]float64
	var a int
	for i:= 0; i <len(num); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&a)
		if a >0 {
		num[i] = math.Sqrt(float64(a))
		}else{
			num[i]=-1
		}
	}
	f.Println(num)
}