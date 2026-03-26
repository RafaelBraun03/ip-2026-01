package main

import f "fmt"

func  main ()  {
	
	var salario, kw, custo_kw, custo_consumo, custo_desconto float32

	  f.Println("Digite o valor do salário mínimo: ")
      f.Scan(&salario)
 
      f.Println("Digite a quantidade de Kw gasto por uma residência: ")
      f.Scan(&kw)

	  custo_kw = salario*0.007
 
      custo_consumo = custo_kw * kw
 
      custo_desconto = custo_consumo*0.9

	f.Printf("Custo po Kw: %.2f", custo_kw)
    f.Println("")
    f.Printf("Custo do consumo: %.2f", custo_consumo)
    f.Println("")
    f.Printf("Custo com desconto: %.2f", custo_desconto)
}