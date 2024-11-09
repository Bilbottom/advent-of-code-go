package main

import (
	"os"
	"testing"
)

var example = ``

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "example", input: example, want: 0},
		{name: "actual", input: input, want: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "actual" && os.Getenv("CI") != "" {
				t.Skip("skipping in CI environment (no `input.data` file)")
			}

			got := part1(test.input)
			if got != test.want {
				t.Errorf("part1() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
		skip  bool
	}{
		{name: "example", input: example, want: 0},
		{name: "actual", input: input, want: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "actual" && os.Getenv("CI") != "" {
				t.Skip("skipping in CI environment (no `input.data` file)")
			}

			got := part2(test.input)
			if got != test.want {
				t.Errorf("part2() = %v, want %v", got, test.want)
			}
		})
	}
}
