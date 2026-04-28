package main

import f "fmt"

func inverter ( arr []int, fim int, inicio int ){
	if inicio >= fim{
		return 
	}
	
	arr[inicio], arr[fim] = arr[fim], arr[inicio]
	 inverter(arr, fim -1, inicio +1)	
}

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n:=len(array)

	f.Println("Original: ", array)
	inverter(array, n-1, 0)
	f.Println("Inversa: ", array)	


}