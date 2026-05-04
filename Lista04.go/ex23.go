package main

import f "fmt"
		
func main (){
	corredor := make([]int, 24)
	janela := make([]int,24)
	var assento int
	for{
		f.Println("Selecione a opção desejada:")
		f.Println("1. Corredor")
		f.Println("2. Janela")
		f.Println("3. Mostrar assentos")
		f.Println("4. Sair")
		var x int
		f.Scan(&x)
		
		switch {
		case x == 1:
			encontrou:=false
			f.Printf("Poltronas disponíveis no corredor:")
			for i,v:= range corredor{
				if v == 0{
					f.Printf("%d, ",i+1)
					encontrou = true
				}
			}
			if !encontrou {
				f.Println("Não há poltronas disponíveis no corredor!")
				f.Println("")
				break
			}
			f.Println("Escolha a Poltrona:")
			f.Scan(&assento)
			if assento < 1 || assento > 24 {
				f.Println("Poltrona inválida")
				f.Println("")
				break
			}
			if corredor[assento-1] == 1{
				f.Println("Assento já ocupado")
				f.Println("")
				break
			}
			 corredor[assento-1] = 1
			 f.Println("Passagem comprada!")
			 f.Println()

		case x == 2:
			encontrou:=false
			f.Printf("Poltronas disponíveis na janela:")
			for i,v:= range janela{
				if v == 0{
					f.Printf("%d, ",i+1)
					encontrou = true
				}
			}
			if !encontrou {
				f.Println("Não há poltronas disponíveis na janela!")
				f.Println("")
				break
			}
			f.Println("Escolha a Poltrona:")
			f.Scan(&assento)
			if assento < 1 || assento > 24 {
				f.Println("Poltrona inválida")
				f.Println("")
				break
			}
			if janela[assento-1] == 1{
				f.Println("Assento já ocupado")
				f.Println("")
				break
			}
			 janela[assento-1] = 1
			 f.Println("Passagem comprada!")
			 f.Println()

		case x==3:
			f.Println("------------------------------------------------------------")
			f.Println("Corredor:",corredor)
			f.Println("Janela:",janela)
			f.Println("------------------------------------------------------------")

		case x == 4:
			return
		}
	}
}