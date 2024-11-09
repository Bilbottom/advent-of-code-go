package main

import (
	"testing"
)

var example1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "example", input: example1, want: 142},
		{name: "actual", input: input, want: 55130},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("part1() = %v, want %v", got, test.want)
			}
		})
	}
}

var example2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "example", input: example2, want: 281},
		{name: "actual", input: input, want: 54985},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := part2(test.input)
			if got != test.want {
				t.Errorf("part2() = %v, want %v", got, test.want)
			}
		})
	}
}