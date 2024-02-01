package matematica

func Operacao[T int | float64](a, b T) T {
	operacao := a + b
	return operacao
}
