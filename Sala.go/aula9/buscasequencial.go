package main

import f "fmt"

func buscasequencial(l []int,x int)int {
	for i:=0; i<=len(l);i++{
		if l[i] == x{
		return i
		}
	}
	return -1 // se o valor não for encontrado retorna -1
}
func main() {
	var lista=[]int{1, 10, 12, 9, 6, 5, 3}
	f.Println(buscasequencial(lista, 6))
}
// A busca sequencial vai rdevolver o indice dentro do array que corresponde ao número procurado