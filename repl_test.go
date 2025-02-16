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
			input: "charzard fire type",
			expected: []string{"charzard", "fire", "type"},
		},
		{
			input: "      pika? squatch     not  (poke) ",
			expected: []string{"pika?", "squatch", "not", "(poke)"},
		},
		{
			input: "THE KING SHOULD EAT FOOD",
			expected: []string{"the", "king", "should", "eat", "food"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected: %v, actual: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v, actual: %v", expectedWord, word)
			}
		}
	}
}