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

func main() {
	var a [fila][col]int
	rand.Seed(time.Now().Unix())

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			a[f][c] = rand.Intn(10)
		}
	}
	fmt.Println("la matriz cargada es: ", a)

	res := sumar(a)
	fmt.Println("la suma de los elementos de la matriz es: ", res)
}

func sumar(a [fila][col]int) int {
	var suma int
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			suma = suma + a[f][c]
		}
	}
	return suma
}
