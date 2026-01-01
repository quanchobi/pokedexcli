package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		// test a normal string
		{
			input: "Hello World",
			expected: []string{"hello", "world"},
		},
		// extra spaces
		{
			input: "   Hello   World   !  ",
			expected: []string{"hello", "world", "!"},
		},
		// empty string
		{
			input: "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("FAIL: actual: %v, expected: %v", actual, c.expected)
		}

		for i:= range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("FAIL: actual word: %s, expected word: %s", actual, c.expected)
			}
		}
	}
}
