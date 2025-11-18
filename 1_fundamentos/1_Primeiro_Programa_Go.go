package main

// Regra 1
// Programas executáveis iniciam pelo pacote main (package main)
// package main --> Programas executaveis em go deve usar o package main  - se for do contrario ira dar erro
// quando for roda use: go run .      ou go build

// Regra 2
// Você NUNCA deve usar package com nome diferente da pasta se a pasta NÃO for módulo independente
/* exemplo:
		pasta: 1_fundamentos
		arquivo: package __fundamentos <-- ERRADO

PACOTES (package) - não precisam ter o mesmo nome da pasta, MAS neste caso NÃO É UM PACOTE, É UM PROGRAMA!
*/

// Regra 3
// Para usar o package como __fundamentos, ou fundamentos, então remove-se o main
/*
Se você realmente quer criar um PACOTE (PACKAGE), então:
- renomeie para package fundamentos
- não coloque a função main()
- importará de outro arquivo

						Criar funções, structs ou constantes que serão usadas por outro arquivo

						Exemplo:
									Arquivo: fundamentos.go

Os códigos em Go são organizados em pacotes e para usá-los é necessário declarar um ou vários (imports) -> Palavra reservada
*/

/*
- interagir com console (é um aplicativo ou interface que permite a interação com um sistema através de entrada e saída de texto, como um terminal de linha de comando etc.)
- debug (exibir informações)
- formatar strings etc.
*/
import "fmt" // biblioteca padrão para formatar e imprimir dados de entrada e saída (I-input/O-output)

// A porta de entrada de um programa Go é a função (main) - Ponto de Partida no APP
func main() {
	fmt.Print("Primeiro Programa Go")
	fmt.Println(" Seja muito bem vindo!")
	fmt.Printf("- \n %s", "Usuario -  Explorando os comandos do printf \\ ou %  ")

	fmt.Println("\n Resumo rápido  - Go suporta vários símbolos de escape tanto para strings quanto para formatadores fmt.Printf" +
		"\n✓ Escapes com barra (\\) → dentro de strings normais funcionam com \\" +
		"\n✓ Formatadores (%) → Não é um escape de string e sim um placeholder de formatação, usados em fmt.Printf() " +
		"\n✓ Strings RAW (` `) → não usam/suportam escapes ELA NÃO INTERPRETA EXEMPLO: \\n")
}

// Para rodar pelo terminal use :
//cd 1_fundamentos
//go run .

/*
	Sobre comentários :
	1 - Priorize código legível e faça comentários que agrega valor!
	2 - Evite comentários óbvios
	3 - Durante o curso abuse dos comentarios
*/
