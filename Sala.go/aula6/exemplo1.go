package main
import "fmt"
func main() {
  const numNotas int = 3 // Quantidade de notas
  var (
	nota [numNotas]float64 // Array de notas
	soma float64 = 0
  )
  // Parte 1 – Leitura
  for i := 0; i < numNotas; i++ {
	fmt.Printf("Informe a nota %d: ", i )
	fmt.Scan(&nota[i]) // Armazena nota no array
	soma += nota[i] // Calcula a soma das notas
  }
  
  for i, v := range nota {
  	  fmt.Printf("Nota %d = %f\n", i, v)
  }
  fmt.Printf("Soma da notas %f: ", soma)
}
