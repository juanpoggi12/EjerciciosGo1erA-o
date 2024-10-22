package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	fila = 5
	col  = 3
)

func main() {
	var A [fila][col]int
	var B [10]int
	var contador, suma int

	rand.Seed(time.Now().Unix())

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			A[f][c] = rand.Intn(10)
		}
	}
	for x := 0; x < 10; x++ {
		B[x] = rand.Intn(10)
	}

	fmt.Println("la matriz A es:", A)
	fmt.Println("el vector B es:", B)

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {

			for x := 0; x < 10; x++ {

				if A[f][c] == B[x] {
					contador = contador + 1
				}

			}
			if contador == 0 {
				suma = suma + A[f][c]
			}
			contador = 0

		}
	}

	fmt.Println("la suma de los elementos de A que no estan en B es:", suma)

}
