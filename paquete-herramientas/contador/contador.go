package contador

import "fmt"

func ContarVocales() {
	var texto string

	fmt.Println("CONTADOR DE VOCALES📖")
	fmt.Print("Ingrese una palabra o texto: ")
	fmt.Scanln(&texto)

	a := 0
	e := 0
	i := 0
	o := 0
	u := 0

	for idx := 0; idx < len(texto); idx++ {
		letra := texto[idx]

		switch letra {
		case 'a', 'A':
			a++
		case 'e', 'E':
			e++
		case 'i', 'I':
			i++
		case 'o', 'O':
			o++
		case 'u', 'U':
			u++
		}
	}

	fmt.Println("Número de veces que aparece cada vocal:")
	fmt.Println("Vocal A:", a)
	fmt.Println("Vocal E:", e)
	fmt.Println("Vocal I:", i)
	fmt.Println("Vocal O:", o)
	fmt.Println("Vocal U:", u)
}
