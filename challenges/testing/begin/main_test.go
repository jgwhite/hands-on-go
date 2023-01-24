// challenges/testing/begin/main_test.go
package main

import "testing"

func TestLetterCount(t *testing.T) {
	counter := letterCounter{}
	cases := map[string]int{
		"#00":       0,
		"a2_32_^_&": 1,
		"812_%6ab/": 2,
	}

	for input, want := range cases {
		t.Run(input, func(t *testing.T) {
			got := counter.count(input)
			if got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}

func TestNumberCount(t *testing.T) {
	counter := numberCounter{}
	cases := map[string]int{
		"#00":        2,
		"abc_1,?/":   1,
		"abc_12&8_^": 3,
	}

	for input, want := range cases {
		t.Run(input, func(t *testing.T) {
			got := counter.count(input)
			if got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}

func TestSymbolCount(t *testing.T) {
	counter := symbolCounter{}
	cases := map[string]int{
		"#00":         1,
		"abc_1,?/":    4,
		"abc_12&8_^_": 5,
	}

	for input, want := range cases {
		t.Run(input, func(t *testing.T) {
			got := counter.count(input)
			if got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}
