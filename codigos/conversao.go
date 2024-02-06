package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var temperaturaKString string
	fmt.Print("Qual a temperatura em K: ")
	fmt.Scanln(&temperaturaKString)

	temperaturaKevin, err := strconv.Atoi(temperaturaKString)
	if err != nil {
		fmt.Println("Valor não válido para conversão para número")
		os.Exit(-1)
	}

	fmt.Printf("A temperatura em Celsius é: %d\n", temperaturaKevin-273)
}
