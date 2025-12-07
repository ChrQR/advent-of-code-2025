package main

import (
	"strconv"
	"testing"
)

func TestCheckData(t *testing.T) {

	tc := []struct {
		data     [][]string
		expected int
	}{
		{
			data: [][]string{
				{"@.@@@."},
				{".@@.@@"},
				{"@.@.@."},
			},
			expected: 7,
		},
	}
	for i, c := range tc {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			movable := checkData(c.data)
			if movable != c.expected {
				t.Errorf("Expected: %d, got: %d", c.expected, movable)
			}
		})
	}

}
