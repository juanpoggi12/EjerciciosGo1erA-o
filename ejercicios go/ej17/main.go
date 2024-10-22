package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	fila = 5
	col  = 5
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
	fmt.Println("la matriz cargada es: ", A)

	for f := 0; f < fila; f++ {
		A = ordenar(A, f)

	}
	fmt.Println("la matriz A con las filas ordenadas es: ", A)

}

func ordenar(A [fila][col]int, f int) [fila][col]int {

	for x1 := 0; x1 < fila-1; x1++ {
		for x2 := x1 + 1; x2 < fila; x2++ {
			if A[f][x1] > A[f][x2] {
				aux := A[f][x1]
				A[f][x1] = A[f][x2]
				A[f][x2] = aux
			}
		}
	}
	return A
}
