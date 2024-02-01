package main

func soma(a *int, b *int) int {
	*a = 60
	return *a + *b
}

func main() {
	var valor1 = 10
	valor2 := 1
	println(soma(&valor1, &valor2))
	println(valor1)

}
