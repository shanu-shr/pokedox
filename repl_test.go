package main

import "testing"

type Case struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {

	cases := []Case{
		{
			input:    " hello    world ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words do not match")
			}
		}
	}
}
