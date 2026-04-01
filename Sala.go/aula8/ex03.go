package main

import f "fmt"

func media(a, b, c float64)float64{
	m := (a +b +c)/3
	return m
}
func main() {
	var a, b, c float64
	f.Println("insira três valores reais:")
	f.Scan(&a, &b, &c)
	f.Println("A média dos três valores é:", media(a, b, c))
}