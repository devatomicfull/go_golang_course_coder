package main

import (
	"fmt"    // pacote de formatação
	m "math" // alias para o pacote math
)

func main() {
	// ======================= CONSTANTES =======================
	const PI float64 = 3.1415 // tipo declarado explicitamente (float64)

	// tipo inferido -> raio será float64 porque PI * float64 exige float64
	var raio = 3.2

	// forma reduzida (somente dentro de funções)
	area := PI * m.Pow(raio, 2)
	fmt.Println("Área da circunferência:", area)

	// bloco de constantes
	const (
		a = 1 // tipo padrão: int
		b = 2 // tipo padrão: int
		c = 3 // tipo padrão: int
	)

	// bloco de variáveis
	var (
		a1 = 1 // tipo inferido: int
		b1 = 2 // tipo inferido: int
		c1 = 3 // tipo inferido: int
	)

	fmt.Println("Constantes:", a, b, c)
	fmt.Println("Variáveis:", a1, b1, c1)

	// ======================= TIPOS BOOLEANOS =======================
	var verdade, falso = true, false // tipo inferido: bool
	fmt.Println("Booleanos:", verdade, falso)

	ligado, desligado := false, true // inferência com :=
	fmt.Println("Ligado/Desligado:", ligado, desligado)

	// ======================= TIPOS NUMÉRICOS =======================

	// inteiros com tipos explícitos
	var i8 int8 = -128
	var ui8 uint8 = 255
	var i32 int32 = -200000
	var ui32 uint32 = 4000000000

	// inteiros sem tipo explícito -> padrão é int
	var inteiro = 10

	// ponto flutuante explícito
	var f32 float32 = 3.14
	var f64 float64 = 6.28 // tipo padrão de float

	// ponto flutuante inferido -> padrão é float64
	flutuante := 9.81

	fmt.Println("Int8:", i8, "Uint8:", ui8)
	fmt.Println("Int32:", i32, "Uint32:", ui32)
	fmt.Println("Int padrão:", inteiro)
	fmt.Println("Float32:", f32, "Float64:", f64, "Inferido:", flutuante)

	// ======================= STRINGS E RUNES =======================

	var texto string = "Hello Go"  // tipo explícito string
	nome := "Inferência de string" // tipo padrão: string

	// rune é um tipo usado para representar caracteres Unicode
	var caractere rune = 'A' // rune = alias para int32

	// complex é um tipo usado para representar números complexos
	simbolo := 'Z' // inferido como rune

	fmt.Println(texto)
	fmt.Println(nome)
	fmt.Println("Rune explícito:", caractere, "Rune inferido:", simbolo)

	// ======================= TIPOS NUMÉRICOS COMPLEXOS =======================
	// uma parte real e uma parte imaginária (na forma a + bi)
	// O tipo complex64 usa componentes float32 para as partes real e imaginária, enquanto complex128 usa componentes float64.
	//Funções integradas como complex(), real() e imag() são usadas para criar e extrair as partes desses números.

	comp := complex(3, 4) // Cria o número complexo 3 + 4i
	var complexo64 complex64 = 1 + 2i
	var complexo128 complex128 = 3 + 4i

	compInferido := 2 + 3i // tipo padrão: complex128

	fmt.Println("Complexo inferido:", c) // // Imprime: (3+4i)
	realPart := real(comp)               // realPart é 3.0
	imagPart := imag(comp)               // imagPart é 4.0

	fmt.Println("Parte real:", realPart)
	fmt.Println("Parte imaginaria:", imagPart)

	fmt.Println("Complex64:", complexo64)
	fmt.Println("Complex128:", complexo128)
	fmt.Println("Complex inferido:", compInferido)

	// ======================= MULTI ATRIBUIÇÃO =======================

	numero, teste, text := 7, false, "Funcionando"

	fmt.Println("Número:", numero)
	fmt.Println("Teste:", teste)
	fmt.Println("Texto:", text)

	// ======================= FORMATANDO NO PRINTF =======================

	fmt.Println("\n")
	fmt.Print("Mesma")
	fmt.Print(" linha.")
	fmt.Println(" Nova")
	fmt.Println("linha.")
	x := 3.141516
	// fmt.Println("O valor de x é " + x)
	// No Go, fmt.Sprint é uma função que converte valores em string sem imprimir nada no console.
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x é %.2f.", x)

	aT := 1
	bT := 1.9999
	cT := false
	dT := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", aT, bT, bT, cT, dT)
	fmt.Printf("\n%v %v %v %v", aT, bT, cT, dT)
}
