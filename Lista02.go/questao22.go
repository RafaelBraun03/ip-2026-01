package main

import f "fmt"

func main() {

	var matricula int

	var h, s_bruto, s_liq float32
	const sal_m = 788.00
	const v_h = 10.00
	
	f.Println("Insira o número de matricula: ")
	f.Scan(&matricula)

	f.Println(" Insira as horas-extras trabalhadas:")
	f.Scan(&h)

	s_bruto = 3*sal_m + v_h*h
	
	if s_bruto > 1500.00 && s_bruto <= 2000.00{
		s_liq = 0.88*s_bruto
		f.Printf("Salário bruto =%.2f", s_bruto)
		f.Println("")
		f.Printf("Salário líquedo =%.2f", s_liq)
	}
	if s_bruto > 2000.00 {
		s_liq = 0.80*s_bruto
		f.Printf("Salário bruto = %.2f", s_bruto)
		f.Println("")
		f.Printf("Salário líquedo =%.2f", s_liq)
	}
}