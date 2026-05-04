package main

import f "fmt"
		
func main (){
	conta := make([]int, 10)
	saldo := make([]float64,10)
	var numero int
	var valor float64
	f.Println("Insira os dados das contas(Número e saldo)")
	s:=0.0
	for i:=0;i<10;i++{
		f.Printf("Conta %d :",i+1)
		f.Scan(&conta[i])
		f.Printf("Saldo:")
		f.Scan(&saldo[i])
		s = s + saldo[i]
	}
	for{
		f.Println("Selecione a opção desejada:")
		f.Println("1. Efetuar depósito")
		f.Println("2. Efetuar saque")
		f.Println("3. Consultar o ativo bancário (ou seja, o somatório dos saldos de todos os clientes)")
		f.Println("4. Finalizar o programa")
		var x int
		f.Scan(&x)
		encontrou:=false
		switch {
		case x == 1:
			f.Println("Insira a conta que receberá o deposito:")
			f.Scan(&numero)
			for i,v:=range conta{
				if numero==v{
					encontrou = true
					f.Println("Saldo atual:R$",saldo[i])
					f.Println("Insira o valor a ser depositado:")
					f.Scan(&valor)
					saldo[i] = saldo[i] + valor
					f.Printf("Saldo atual da conta %d: R$%.2f",numero,saldo[i])
					f.Println("")
				}
				
			}
				if !encontrou{
					f.Println("Conta não encontrada")
					f.Println("")
					break
				}
		case x == 2:
			f.Println("Insira a conta que será feito o saque:")
			f.Scan(&numero)
			for i,v:=range conta{
				if numero==v{
					encontrou = true
					f.Println("Saldo atual:R$",saldo[i])
					f.Println("Insira o valor a ser sacado:")
					f.Scan(&valor)
					if valor <= saldo[i]{
					saldo[i] = saldo[i] - valor
					f.Println("Saque realizado!")
					f.Printf("Saldo atual da conta %d: R$%.2f",numero,saldo[i])
					f.Println("")
				}else{
					f.Println("Saldo insuficiente")
					break
				}
				}
				}
				if !encontrou{
					f.Println("Conta não encontrada")
					f.Println("")
					break
				}
			
		case x == 3:
			   s = 0
   			 for _, v := range saldo {
        		s += v
    		}
			f.Printf("Ativo bancário:R$%.2f",s)
			f.Println("")
			break
		case x == 4:
			return
		}
	}		
}
