package main

import (f "fmt"
		 "strconv")


func converter(n int)string{
	if n == 0{
		return "0"
	}
	if n ==1 {
		return  "1"
	}
	return  converter(n/2) + strconv.Itoa(n%2)
}
func main()  {
	var n int
	f.Println("Insira um número:")
	f.Scan(&n)
	if n <0 {
		f.Println("Número inválido")
		return 
	}
	f.Printf("%d em binário é: %s", n, converter(n))
}