package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.data
var input string

func main() {
	fmt.Println("--- year __ day __ ---")
	data := strings.TrimSuffix(input, "\n")
	fmt.Println("    part 1:", part1(data))
	fmt.Println("    part 2:", part2(data))
}

func part1(data string) any {

	return ""
}

func part2(data string) any {

	return ""
}
