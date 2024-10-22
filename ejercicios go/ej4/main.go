package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	fila = 3
	col  = 3
)

func main() {
	var a [fila][col]int
	var num int
	rand.Seed(time.Now().Unix())
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			a[f][c] = rand.Intn(5)
		}
	}
	fmt.Println("la matriz cargada es: ", a)

	fmt.Println("ingrese N")

	fmt.Scanln(&num)

	res := numero(a, num)
	fmt.Println("la cantidad de veces que aparece N en la matriz es: ", res)
}

func numero(a [fila][col]int, num int) int {
	var contador int

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			if a[f][c] == num {
				contador = contador + 1
			}
		}
	}
	return contador
}
