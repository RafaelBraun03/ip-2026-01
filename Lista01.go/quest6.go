package main

import f "fmt"

func main () {

	var  f1, f2, f3, c1, c2, c3  float32

	f.Println("Digite a temperatura em  Fahrenheit : ")
    f.Scan(&f1)
	f.Println("Digite a temperatura em  Fahrenheit : ")
    f.Scan(&f2)
	f.Println("Digite a temperatura em  Fahrenheit : ")
    f.Scan(&f3)

	 c1 = (5*(f1 - 32))/9
  
     c2 = (5*(f2 - 32))/9
  
     c3 = (5*(f3 - 32))/9

	f.Println(f1, "FAHRENHEIT EQUIVALE A ", c1, " ", "CELSIUS")
    f.Println(f2, "FAHRENHEIT EQUIVALE A ", c2, " ", "CELSIUS")
    f.Println(f3, "FAHRENHEIT EQUIVALE A ", c3, " ", "CELSIUS")
}