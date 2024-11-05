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
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HeLlO wOrLd ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)

		if len(result) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", result, c.expected)
		}

		for i := range result {
			word := result[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, result, c.expected)
			}
		}
	}
}
