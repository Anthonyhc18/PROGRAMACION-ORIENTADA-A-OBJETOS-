package main

import (
	"fmt"
	"paquete-herramientas/contador"
	"paquete-herramientas/conversor"
)

func main() {
	var opcion string

	for {
		fmt.Println("MENU PRINCIPAL")
		fmt.Println("1. Conversor de monedas")
		fmt.Println("2. Contador de vocales")
		fmt.Println("3. Salir")
		fmt.Print("Elija una opción (1-3): ")
		fmt.Scan(&opcion)

		if opcion == "1" {
			conversor.Convertir()
		} else if opcion == "2" {
			contador.ContarVocales()
		} else if opcion == "3" {
			fmt.Println("ADIOS")
			break
		} else {
			fmt.Println("Opción incorrecta, intente de nuevo.")
		}
	}
}
