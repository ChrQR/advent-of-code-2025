package main

import (
	"testing"
)

func TestMaxJolt(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{
			Input:    "3233434223352253322244323562413222322422522622312422332123223123235422212196323222332232332242222211",
			Expected: 963342222211,
		},
	}

	for _, c := range cases {
		t.Run("test TestMaxJolt", func(t *testing.T) {
			num := highestNumber(c.Input)
			if num != c.Expected {
				t.Log(c.Expected > num)
				t.Errorf("expected: %d, got %d", c.Expected, num)
			}
		})
	}
}
