package main

import "fmt"

func main() {
	// The make function allocates and initializes a hash map data
	// structure and returns a map value that points to it.
	map1 := make(map[int]string)
	fmt.Println(map1)

	// With make we can specify the inital size of the map
	map2 := make(map[int]string, 20)
	fmt.Println(map2)

	// The same goes for composite literals
	map3 := map[int]string{}
	fmt.Println(map3)

	// With composite literals we can directly
	// initalize the map with specific values
	players := map[int]string{
		10: "Totti",
		16: "de Rossi",
		24: "Florenzi",
	}
	fmt.Println(players)
}
