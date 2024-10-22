package main

import (
	"fmt"
	"math/rand"
	"time"
)

const fila = 6
const col = 5

func main() {
	var a [fila][col]int
	rand.Seed(time.Now().Unix())
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			a[f][c] = rand.Intn(10)
		}
	}
	fmt.Println("la matriz cargada es: ", a)

	res1, res2 := sumar(a)
	fmt.Println("la suma de los elementos de la matriz mayores o iguales a 5 es: ", res1)
	fmt.Println("la cantidad de elementos que intervinieron es: ", res2)

}

func sumar(a [fila][col]int) (int, int) {
	var suma, contador int
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			if a[f][c] >= 5 {
				suma = suma + a[f][c]
				contador = contador + 1
			}
		}
	}
	return suma, contador
}
