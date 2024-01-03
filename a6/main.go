package main

import "fmt"

const a = "bicicleta"

type id int

var (
	b bool
	c int
	e float64
	f string
	d id = 12
)

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s), s[:4])

	outro := s[:4]
	fmt.Println(outro)
}
