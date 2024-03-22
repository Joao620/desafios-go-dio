package main

import (
	"fmt"
)

var operacoes = map[string]func(int, int) int{
    "mul": func(a int, b int) int { return a * b },
    "div": func(a int, b int) int { return a / b },
	"add": func(a int, b int) int { return a + b },
	"sub": func(a int, b int) int { return a - b }
}

func main() {
	fmt.Println("Qual operação deseja realizar? (add, sub, mul, div)")
	var op string
	fmt.Scanln(&op)
	fmt.Println("Digite os numeros separados por um espaço (numeros inteiros only!)")
	var a, b int
	fmt.Scanln(&a, &b)

	if operacao, ok := operacoes[op]; ok {
		if op == "div" && b == 0 {
			fmt.Println("Erro: Divisão por zero não é permitida.")
		} else {
			result := operacao(a, b)
			fmt.Printf("Resultado: %d\n", result)
		}
	} else {
		fmt.Println("Operação inválida")
	}
}
