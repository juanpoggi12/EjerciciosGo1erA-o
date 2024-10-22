package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	fila = 6
	col  = 6
)

func main() {
	var A [fila][col]int
	rand.Seed(time.Now().Unix())

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			if c == 2 {
				A[f][c] = -6
			}
			A[f][c] = rand.Intn(5)
		}
	}

	fmt.Println("la matriz A es:", A)

	res := cambiar(A)
	fmt.Println("la matriz cambiada es: ", res)

}
func cambiar(A [fila][col]int) [fila][col]int {
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			if A[f][c] < 0 {
				A[f][c] = 0
			} else {
				if A[f][c] == 0 {
					A[f][c] = 1
				} else {
					A[f][c] = 2
				}
			}
		}
	}

	return A
}
