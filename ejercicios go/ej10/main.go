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

	fmt.Println("la matriz A cargada es:", A)

	res1 := diagonal(A)
	fmt.Println("la suma de los elementos de la diagonal principal es: ", res1)

}

func diagonal(A [fila][col]int) int {
	var suma int

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			if f == 0 && c == 0 {
				suma = suma + A[f][c]
			} else {
				if f == 1 && c == 1 {
					suma = suma + A[f][c]
				} else {
					if f == 2 && c == 2 {
						suma = suma + A[f][c]
					} else {
						if f == 3 && c == 3 {
							suma = suma + A[f][c]
						}
					}
				}
			}
		}
	}

	return suma
}
