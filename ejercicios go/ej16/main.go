package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	fila = 4
	col  = 4
)

var (
	A [fila][col]int
)

func main() {
	rand.Seed(time.Now().Unix())
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			A[f][c] = rand.Intn(10)
		}
	}
	fmt.Println("la matriz A es: ", A)

	res1 := cambiar(A)
	fmt.Println("la matriz cambiada es: ", res1)
}

func cambiar(A [fila][col]int) [fila][col]int {

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {

			if f == 0 {
				aux := A[f+1][c]
				A[f+1][c] = A[f][c]
				A[f][c] = aux
			} else {
				if f == 2 {
					aux := A[f+1][c]
					A[f+1][c] = A[f][c]
					A[f][c] = aux
				}
			}

		}
	}
	return A
}
