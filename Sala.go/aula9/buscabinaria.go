package main

import f "fmt"

func buscabinaria(l []int, x int)int{ 
	esquerda:=0
	direita:=len(l)
	for esquerda<=direita{
		meio:=(esquerda+ direita)/2
		elemento_meio:=l[meio]
		if elemento_meio == x{
			return  meio
		}
		if elemento_meio > x{
			direita = meio-1
		}
		if elemento_meio < x {
			esquerda = meio + 1
		}
	}
	return -1
}
func main() {
	var lista=[]int{1, 3, 5, 6, 8, 11, 14, 15}
	f.Println(buscabinaria(lista, 6))
}