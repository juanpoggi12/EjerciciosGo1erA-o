package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	fila = 4
	col  = 5
)

func main() {
	var A [fila][col]int
	rand.Seed(time.Now().Unix())

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			A[f][c] = rand.Intn(10)
		}
	}

	fmt.Println("la matriz cargada es: ", A)
	res1 := filas(A)
	fmt.Println("el vector con las sumas de los elementos de cada fila de A es: ", res1)
	res2 := columnas(A)
	fmt.Println("el vector con las sumas de los elementos de cada columna de A es: ", res2)
}

func filas(A [fila][col]int) [4]int {
	var suma int
	var vector [4]int
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			suma = suma + A[f][c]
		}
		vector[f] = suma
		suma = 0
	}

	return vector
}

func columnas(A [fila][col]int) [col]int {
	var suma int
	var vector [col]int

	for c := 0; c < col; c++ {
		for f := 0; f < fila; f++ {
			suma = suma + A[f][c]
		}
		vector[c] = suma
		suma = 0
	}
	return vector
}
