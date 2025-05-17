package main

import (
	"testing"
)

func TestLuhn(t *testing.T) {
	type Test struct {
		cardNum string
		valid   bool
	}

	tests := []Test{
		//	fill slice with cards
		Test{"1234567890123456", false},
		Test{"4408041234567893", true},
		Test{"38520000023237", true},
		Test{"4222222222222", true},
	}

	for _, test := range tests {
		res := luhn(test.cardNum)
		if res != test.valid {
			t.Errorf("luhn(%q) = %v; want %v", test.cardNum, res, test.valid)
		}
	}
}
