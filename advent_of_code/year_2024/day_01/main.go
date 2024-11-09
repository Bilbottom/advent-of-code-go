package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.data
var input string

func main() {
	fmt.Println("--- year 2024 day 01 ---")
	data := strings.TrimSuffix(input, "\n")
	fmt.Println("    part 1:", part1(data))
	fmt.Println("    part 2:", part2(data))
}

func part1(data string) any {

	return 0
}

func part2(data string) any {

	return 0
}
