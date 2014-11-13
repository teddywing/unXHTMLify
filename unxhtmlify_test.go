package main

import "testing"

func TestUnxhtmlifyLine(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{ "<br />", "<br>" },
	}
	
	for _, c := range tests {
		result := unxhtmlify_line(c.input)
		if result != c.expected {
			t.Errorf("unminify_line(%q) == %q, want %q", c.input, result, c.expected)
		}
	}
}
