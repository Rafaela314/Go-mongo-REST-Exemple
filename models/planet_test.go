package models

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		planet   Planet
		expected bool
	}{
		{planet: Planet{Name: "", Weather: "sunny", Land: 0, Appearances: 1}, expected: false},
		{planet: Planet{Name: "Tatooine", Weather: "sunny", Land: 4789.345, Appearances: 1}, expected: true},
	}

	for _, test := range tests {
		if test.planet.IsValid() != test.expected {
			t.Errorf("%v: got %t, want %t", test, test.planet.IsValid(), test.expected)
		}
	}
}
