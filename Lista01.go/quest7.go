package main

import f "fmt"

func main () {

	var f1, c1, chuva, volume float32

	 f.Println("Digite a temperatura em  Fahrenheit : ")
     f.Scan(&f1)

     f.Println("Digite a quantidade de chuva em polegadas: ")
     f.Scan(&chuva)

	 volume = chuva*25.4
	 c1 = (5*(f1 - 32))/9

	 f.Printf("O valor em celsius é = %.2f", c1)
	 f.Println("")
     f.Printf("A quantidade de chuva é = %.2f", volume)

}