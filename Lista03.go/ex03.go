package main

import f "fmt"
        

func main() {
	var sal_carlos float64
	var meses int

	f.Println("Insira o valor do salário do Carlos:")
	f.Scan(&sal_carlos)
	sal_joao := sal_carlos/3
  
	total_C := sal_carlos 
	total_j := sal_joao
// consideraremos que não haverá aportes ao longo dos meses e que o valor 
for total_j < total_C {
	total_C = total_C*1.02
	total_j = total_j*1.05
// eles ainda recebem salário ao longo dos meses, devemos acresentar
	total_C= total_C+ sal_carlos
	total_j = total_j + sal_joao

	meses = meses + 1
}
f.Println("A quantidade de meses para João superar Carlos é:", meses)

}