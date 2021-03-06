package main

import "testing"

func TestUnxhtmlifyString(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{ "<br />", "<br>" },
		{ "<span>Bender / Flexo</span>", "<span>Bender / Flexo</span>" },
		{ "<img src=\"farnsworth-paradox.jpg\" alt=\"Parabox\"/>", "<img src=\"farnsworth-paradox.jpg\" alt=\"Parabox\">" },
		{ "some-class-name\" />", "some-class-name\">" },
		{ "/>", ">" },
		{ "		/>", "		>" },
	}
	
	for _, c := range tests {
		result := unxhtmlify_line(c.input)
		if result != c.expected {
			t.Errorf("unminify_line(%q) == %q, want %q", c.input, result, c.expected)
		}
	}
}
