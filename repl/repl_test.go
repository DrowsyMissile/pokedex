package repl

import (
	"testing"
)

func Test_CleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " can   you  split     me?",
			expected: []string{"can", "you", "split", "me?"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("input: %s did not match expected %v, instead got %v", c.input, c.expected, actual)
			continue
		}

		for i, word := range actual {
			expected := c.expected[i]
			if word != expected {
				t.Errorf("word: %s did not match expected %s", word, expected)
				t.Fail()
			}
		}
	}
}
