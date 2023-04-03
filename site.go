package main

import "fmt"

func main() {
	exibiSaudacoes()
	exibiMenu()
	exibiOpcoes()
}

func exibiSaudacoes() {
	nome := "Gustavo Souza"
	fmt.Println("Olá sr.", nome)
	fmt.Println("Escolha uma das opções abaixo")
}

func exibiMenu() {
	fmt.Println("1 - Iniciar Monitoramento:")
	fmt.Println("2 - Exibir Logs:")
	fmt.Println("0 - Sair do Programa:")
}

func leOpcao() int {
	var opcaoLida int
	fmt.Scan(&opcaoLida)
	fmt.Println("A opção escolhida foi: ", opcaoLida)
	return opcaoLida
}
func exibiOpcoes() {
	comando := leOpcao()

	switch comando {
	case 0:
		fmt.Println("Saindo do programa...")
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	default:
		fmt.Println("Não existe essa opção")
	}
}
