package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "hello world !!!!"
	showInteface(x)
	showInteface(y)
}

func showInteface(x interface{}) {

	fmt.Printf("O tipo é %T e valor é %v \n", x, x)
}
