package main

import (
	"fmt"
)

const (
	coldatos = 6
	filares  = 6
	colres   = 7
	claseave = 4
)

var (
	claseanimal, cliente, servicio, importe, atencion, nroservicio int
	datos                                                          [][coldatos]int
	vec                                                            [coldatos]int
	res                                                            [filares][colres]int
	sel[][3]int
)

func main() {

	fmt.Println("ingrese el codigo de clase de animal")
	fmt.Scanln(&claseanimal)

	nroservicio = 999
	///carga de datos
	for claseanimal != 0 {

		fmt.Println("ingrese el codigo de cliente")
		fmt.Scanln(&cliente)
		fmt.Println("ingrese el codigo del servicio ofrecido")
		fmt.Scanln(&servicio)
		fmt.Println("ingrese el importe cobrado")
		fmt.Scanln(&importe)
		fmt.Println("ingrese el tipo de atencion")
		fmt.Scanln(&atencion)

		nroservicio=nroservicio+1

		vec[0] = nroservicio
		vec[1] = cliente
		vec[2] = claseanimal
		vec[3] = servicio
		vec[4] = importe
		vec[5] = atencion

		datos = append(datos, vec)

		res[servicio-1][claseanimal-1] = res[servicio-1][claseanimal-1] + 1 //carga de datos en matriz res

		fmt.Println("ingrese el codigo de clase de animal")
		fmt.Scanln(&claseanimal)

	} //fin del for (carga de datos)
	fmt.Println("la matriz datos es: ", datos)

	c(&res) //totales de matriz res
	fmt.Println("la matriz res es:", res)

	fmt.Println("el nombre sel servicio que mas se atendio para la clase de animal ave es:", d(res, claseave)) //funcion punto d

	e(res)//func punto e

	sel=f(datos)
	fmt.Println("la matriz sel es: ", sel)//funcion punto f

	g(sel)
	for f:=0; f<len(sel); f++{
	fmt.Println("la matriz sel ordenada es: ", sel[f])
	}
	} ///fin main

func c(res *[filares][colres]int) {
	var suma int

	for f := 0; f < filares-1; f++ {
		for c := 0; c < colres-1; c++ {
			suma = suma + res[f][c]
		}
		res[f][colres-1] = suma
		suma = 0
	}

	for c := 0; c < colres; c++ {
		for f := 0; f < filares-1; f++ {
			suma = suma + res[f][c]
		}
		res[filares-1][c] = suma
		suma = 0
	}
}

func d(res [filares][colres]int, codclaseanimal int) string {
	var serviciomayor, valorservicio int
	var resultado string
	for f := 0; f < filares-1; f++ {
		if f == 0 {
			serviciomayor = f + 1
			valorservicio = res[f][codclaseanimal-1]
		} else if res[f][codclaseanimal-1] > valorservicio {
			serviciomayor = f + 1
			valorservicio = res[f][codclaseanimal-1]
		}
	}

	switch serviciomayor {
	case 1:
		resultado = "consulta"
	case 2:
		resultado = "clinico"
	case 3:
		resultado = "preventivo"
	case 4:
		resultado = "etologia"
	case 5:
		resultado = "vacunacion"
	}
	return resultado
}

func e(res[filares][colres]int){
	var porcentaje float32
	var clase string

		for c:=0; c<colres-1; c++{
			porcentaje=float32(res[filares-1][c])*100/float32(res[filares-1][colres-1])

			switch c+1{
			case 1: clase="perros"
			case 2: clase="gatos"
			case 3: clase="conejos"
			case 4: clase="aves"
			case 5: clase="roedores"
			case 6: clase= "otras especies"
			}
			fmt.Println("el porcentaje que la clase de animal", clase, "representa sobre el total", res[filares-1][colres-1], "es:", porcentaje)

		}

}

func f(datos[][coldatos]int)([][3]int){
	var vec[3]int
	var matriz[][3]int

	for f:=0; f<len(datos); f++{
		if datos[f][coldatos-1]==2{
			vec[0]=datos[f][0]
			vec[1]=datos[f][1]
			vec[2]=datos[f][4]

			matriz = append(matriz, vec)
		}
	}
return matriz
}

func g(sel[][3]int){

	for x1:=0; x1<len(sel)-1; x1++{
		for x2:=x1+1; x2<len(sel); x2++{
			if sel[x1][2]>sel[x2][2]{
				for c:=0; c<len(sel[0]); c++{
					aux:=sel[x1][c]
					sel[x1][c]=sel[x2][c]
					sel[x2][c]=aux
				}
			}
		}

	}

}