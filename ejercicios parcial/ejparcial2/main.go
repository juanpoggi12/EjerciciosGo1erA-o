package main

import (
	"fmt"
)

const ()

var (
	telefonos                                                    [][3]int
	llamadas                                                     [][3]int
	nrotel, empresa, minutoslibres, origen, destino, cantminutos, empresafuncC int
	vec                                                          [3]int
)

func main() {

	fmt.Println("ingrese el numero de telefono")
	fmt.Scanln(&nrotel)

	for nrotel != 0 {
		fmt.Println("ingrese la empresa")
		fmt.Scanln(&empresa)
		fmt.Println("ingreselos minutos libres incluidos en el abono de la empresa")
		fmt.Scanln(&minutoslibres)

		vec[0] = nrotel
		vec[1] = empresa
		vec[2] = minutoslibres

		telefonos = append(telefonos, vec)

		fmt.Println("ingrese el numero de telefono")
		fmt.Scanln(&nrotel)
	}

	fmt.Println("ingrese el numero de origen de la llamada")
	fmt.Scanln(&origen)

	for origen != 0 {
		fmt.Println("ingrese el numero de destino")
		fmt.Scanln(&destino)
		fmt.Println("ingrese la cantidad total de minutos de la llamada")
		fmt.Scanln(&cantminutos)

		vec[0] = origen
		vec[1] = destino
		vec[2] = cantminutos

		llamadas = append(llamadas, vec)

		fmt.Println("ingrese el numero de origen de la llamada")
		fmt.Scanln(&origen)

	}



}//fin de main

func c(empresa int){
var minconsumidos int
	for f:=0; f<len(llamadas); f++{
		if empresa==llamadas[f][1]{
			if llamadas[f][0]==telefonos[f][0]{
				minconsumidos=minconsumidos+llamadas[f][2]
			}

		}



	}
}