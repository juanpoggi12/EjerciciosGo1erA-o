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
			a[f][c] = rand.Intn(5)
		}
	}
	fmt.Println("la matriz es: ", a)

	res1, res2 := paridad(a)
	fmt.Println("la suma de los pares es", res1)
	fmt.Println("la suma de los impares es", res2)

}
func paridad(a [fila][col]int) (int, int) {
	var sumapar, sumaimpar int
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			if a[f][c]%2 == 0 {
				sumapar = sumapar + a[f][c]
			} else {
				sumaimpar = sumaimpar + a[f][c]
			}
		}
	}
	return sumapar, sumaimpar
}
