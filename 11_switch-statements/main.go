package main

import "fmt"

func main() {
	// break isn't needed like in other languages
	switch "Totti" {
	case "Strootman":
		fmt.Println("Kevin Strootman")
	case "Taddei":
		fmt.Println("Rodrigo Taddei")
	case "Totti":
		fmt.Println("Francesco Totti")
	case "de Rossi":
		fmt.Println("Daniele de Rossi")
	case "Florenzi":
		fmt.Println("Alessandro Florenzi")
	case "Nainggolan":
		fmt.Println("Radja Nainggolan")
	default:
		fmt.Println("Unknown player")
	}

	// A case body breaks automatically, unless it ends with a fallthrough statement.
	switch "Totti" {
	case "Strootman":
		fmt.Println("Kevin Strootman")
	case "Taddei":
		fmt.Println("Rodrigo Taddei")
	case "Totti":
		fmt.Println("Francesco Totti")
		fallthrough
	case "de Rossi":
		fmt.Println("Daniele de Rossi")
		fallthrough
	case "Florenzi":
		fmt.Println("Alessandro Florenzi")
	case "Nainggolan":
		fmt.Println("Radja Nainggolan")
	}

	// Multiple evals are allowed
	switch "Totti" {
	case "Strootman", "Taddei":
		fmt.Println("Kevin Strootman, Rodrigo Taddei")
	case "Totti", "de Rossi":
		fmt.Println("Francesco Totti, Daniele de Rossi")
	case "Florenzi", "Nainggolan":
		fmt.Println("Alessandro Florenzi, Radja Nainggolan")
	default:
		fmt.Println("Unknown player")
	}

	// if no expression provided, go checks for the first case that evals to true
	player := "Totti"
	switch {
	case player == "Strootman":
		fmt.Println("Kevin Strootman")
	case player == "Taddei":
		fmt.Println("Rodrigo Taddei")
	case player == "Totti":
		fmt.Println("Francesco Totti")
	case player == "de Rossi":
		fmt.Println("Daniele de Rossi")
	case player == "Florenzi":
		fmt.Println("Alessandro Florenzi")
	case player == "Nainggolan":
		fmt.Println("Radja Nainggolan")
	default:
		fmt.Println("Unknown player")
	}
}
