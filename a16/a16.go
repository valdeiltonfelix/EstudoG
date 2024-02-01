package main

import "fmt"

func main() {
	var MinhaInterface interface{} = "Valdeilton de souza felix"
	println(MinhaInterface.(string))
	resultado, ok := MinhaInterface.(int)
	fmt.Println(resultado, ok)
}
