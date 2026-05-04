package main

import f "fmt"

func main() {
	num := make([]int,0,100)
	
	for i:= 1; i <=100; i++ {
		num = append(num, i)
	}
	f.Println(num)
}