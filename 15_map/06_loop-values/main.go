package main

import "fmt"

func main() {
	players := map[int]string{
		10: "Totti",
		16: "de Rossi",
		24: "Florenzi",
	}

	for k, v := range players {
		fmt.Println("Key:", k, "Value:", v)
	}
}
