package main

import (
	"fmt"
	"os"
	"regexp"
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

	fmt.Println("El numero insertado es correcto: ", binaryCode)

}
