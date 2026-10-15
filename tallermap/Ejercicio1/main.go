package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "hola mundo golang hola mundo programa hola"
	palabras := strings.Fields(texto)
	conteo := make(map[string]int)

	for i := 0; i < len(palabras); i++ {
		palabra := palabras[i]
		conteo[palabra] = conteo[palabra] + 1
	}

	palabraMasRepetida := ""
	maxRepeticiones := 0

	for palabra, veces := range conteo {
		fmt.Println(palabra, ":", veces)
		if veces > maxRepeticiones {
			maxRepeticiones = veces
			palabraMasRepetida = palabra
		}
	}

	fmt.Println("La palabra que más se repite es:", palabraMasRepetida, "con", maxRepeticiones, "veces")
}
