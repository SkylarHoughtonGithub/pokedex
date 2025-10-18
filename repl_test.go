package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"  hello   world  ", []string{"hello", "world"}},
		{"Charmander Bulbasaur PIKACHU", []string{"charmander", "bulbasaur", "pikachu"}},
		{"\n\t spaced\tout \n", []string{"spaced", "out"}},
		{"", []string{}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Fatalf("length mismatch: got %d, want %d (input: %q)", len(actual), len(c.expected), c.input)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word %d mismatch: got %q, want %q (input: %q)", i, word, expectedWord, c.input)
			}
		}
	}
}
