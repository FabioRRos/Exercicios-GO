package main

import "fmt"

var num1 int
var num2 int
var result float64
var opcao int

func main() {

	fmt.Println("Bem vindo a calculadora!")
	fmt.Println("1 - Somar")
	fmt.Println("2 - Subitrair")
	fmt.Println("3 - Dividir")
	fmt.Println("4 - Multiplicar")
	_, err := fmt.Scan(&opcao)

	if err == nil && opcao <= 4 {
		fmt.Println("Digite o Primeiro valor")
		fmt.Scan(&num1)
		fmt.Println("Digite o Segundo valor")
		fmt.Scan(&num2)
	}

	switch opcao {
	case 1:
		{
			result = float64(num1) + float64(num2)
		}
	case 2:
		{
			result = float64(num1) - float64(num2)
		}
	case 3:
		{
			result = float64(num1) / float64(num2)
		}
	case 4:
		{
			result = float64(num1) * float64(num2)
		}
	default:
		{
			fmt.Println("Opcao Invalida")
		}
	}

	fmt.Printf("O valor final é: %.2f", result)
}
