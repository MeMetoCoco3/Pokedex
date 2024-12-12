package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "SingleWord!",
			expected: []string{"singleword!"},
		},
		{
			input:    "Pikachu, Pokemon, jj",
			expected: []string{"pikachu,", "pokemon,", "jj"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("We expected %s, and we recieved %s\n", expectedWord, word)
			}
		}
	}
}
