package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var binaryCode string
	matched := false
	var err error

	fmt.Println("Introduce un numero binario de 8 digitos:")

	for !matched {
		fmt.Scanf("%v\n", &binaryCode)
		matched, err = regexp.MatchString(`([0-1]{8})`, binaryCode)

		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		if !matched {
			fmt.Println("El numero insertado no reune los requisitos (8 digitos)")
			fmt.Println("Porfavor inserta un nuevo numero binario:")
		}
	}

	decimal, err := strconv.ParseInt(binaryCode, 2, 64)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println("-------------------------------")
	fmt.Println("El numero insertado es correcto")
	fmt.Println("-------------------------------")
	fmt.Println("Numero binario: ", binaryCode)
	fmt.Println("Numero decimal: ", decimal)
}
