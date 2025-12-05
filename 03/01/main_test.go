package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestMaxJolt(t *testing.T) {
	cases := []struct {
		Input    []string
		Expected int
	}{
		{
			Input: []string{
				"3233434223",
				"3253614853",
				"9227239868",
			},
			Expected: 44 + 85 + 99,
		},
	}

	for _, c := range cases {
		t.Run("test TestMaxJolt", func(t *testing.T) {
			input := strings.Join(c.Input, "\n")
			scanner := bufio.NewScanner(strings.NewReader(input))
			total := MaxJolt(scanner)
			if total != c.Expected {
				t.Errorf("expected: %d, got %d", c.Expected, total)
			}
		})
	}
}
