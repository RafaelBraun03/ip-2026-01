package main

import f "fmt"

func main() {
	var idade, continuar int
	var altura, peso float64

	var totalPessoas, qtdMais50, qtdIdade10a20, qtdPesoMenos40 int
	var somaAltura10a20 float64

	for {
		f.Println("Digite a idade:")
		f.Scan(&idade)
		f.Println("Digite a altura:")
		f.Scan(&altura)
		f.Println("Digite o peso:")
		f.Scan(&peso)
		totalPessoas = totalPessoas + 1

		if idade >50{
			qtdMais50 = qtdMais50 +1
		}
		if idade >=10 && idade <= 20 {
			somaAltura10a20 = somaAltura10a20 + altura
			qtdIdade10a20 = qtdIdade10a20 +1
		}
		
		if peso <40.0 {
			qtdPesoMenos40 = qtdPesoMenos40 + 1
		}
		f.Println("Deseja continuar inserindo dados?(sim=1)")
		f.Scan(&continuar)			

		if continuar !=1 {
			break
		}
	}

	f.Println("Quantidade de pessoas com mais de 50 anos:", qtdMais50)
	
	if qtdIdade10a20 > 0 {
		media:= somaAltura10a20/float64(qtdIdade10a20)
		f.Println("A média das alturas das pessoas de idade entre 10 a 20 anos é", media, "cm")
	}else {
		f.Println("Não há pessoas com idade entre 10 e 20 anos")
	}
	if qtdPesoMenos40>0 {
		porcent:=qtdPesoMenos40/totalPessoas
		f.Println("A porcentagem de pessoas com peso menor que 40 é:", porcent)
	}
}