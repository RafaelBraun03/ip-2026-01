package main

import f "fmt"
       
func maior(a, b, c int)int{
	var m int
     
	 if a > b && a > c{
		m = a
	 }
	 if b > a && b > c {
		m =b
	 }
	 if c > a && c > b {
		m =c
	 }
			
	return m
		 
}
func main() {
	var a, b, c int

	f.Println("Insira três números:")
    f.Scan(&a, &b, &c)

	f.Println("O maior é:", maior(a, b, c))
}