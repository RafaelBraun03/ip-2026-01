package main

import f "fmt"

func main () {

	var a, b, c int

	f.Println("Insira o primeiro valor:")
	f.Scan(&a)
	f.Println("Insira o segundo valor:")
	f.Scan(&b)
	f.Println("Insira o terceiro valor:")
	f.Scan(&c)

	if a > b && a > c {
		if b > c {
			f.Println(c, ",", b, ",", a)
		}else {
			f.Println(b, ",", c, ",", a)
		}
	}
	if b > a && b > c {
		if a > c {
			f.Println(c, ",", a, ",", b)
		}else {
			f.Println(a, ",", c, ",", b)
		}
	}
	
	if c > a && c > b {
		if a > b {
			f.Println(b, ",", a, ",", c)
		}else {
			f.Println(a, ",", b, ",", c)
		}
	}

}