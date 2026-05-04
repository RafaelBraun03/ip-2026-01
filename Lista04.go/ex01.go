package main

import f "fmt"

func main() {
	var num [10]int

	for i:= 0; i <len(num); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&num[i])
	}
	for i,v:= range num {
		if v > 50 {
			f.Printf("Ítem de índice %d = %d; É maior.", i, v)
			f.Println()
		}
	}
	f.Println("Não há números maior que 50")
}