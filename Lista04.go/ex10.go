package main

import f "fmt"

func main() {
	fibon := make([]int,50)
	fibon[0]=1
	fibon[1]=1
	for i:=2;i<50;i++{
		fibon[i]= fibon[i-2]+fibon[i-1]
		fibon = append(fibon, fibon[i])
	}
	f.Print(fibon)
	
}