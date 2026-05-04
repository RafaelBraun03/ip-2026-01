package main

import f "fmt"

func main() {
	num := make([]int,0,100)
	
	for i:= 1; i <=200; i++ {
		if i%2 != 0{
		num = append(num, i)
	}
	}
	f.Println(num)
}