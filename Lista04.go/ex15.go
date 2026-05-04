package main
import f "fmt"

func main() {
	var vetor [30]int
	v2 := make([]int,0,30)
	for i:= 0; i <len(vetor); i++ {
		f.Printf("Digite o %dº valor do vetor: ", i+1)
		f.Scan(&vetor[i])
	}
	for i,v := range vetor{
		if (i+1)%2==0 {    // fazendo para a posição e não para o índice.
			v2 = append(v2, v*2)
		}else{
			v2 = append(v2, v*3)
		}
	}
	f.Println("")
	f.Println("Inicial:",vetor)
	f.Println("-------------------------------")
	f.Println("Final:",v2)
}