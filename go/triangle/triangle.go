//+build !example

package triangle

import (
	"math"
)

const (
	Equ = iota // equilateral
	Iso        // isosceles
	Sca        // scalene
	NaT        // not a triangle
)

type Kind int

func KindFromSides(a, b, c float64) Kind {
	// check for triangle inequality
	if a+b <= c || a+c <= b || b+c <= a {
		return NaT
	}

	switch {
	// edge cases
	case areNaN(a, b, c): // NaN
		fallthrough
	case areInf(a, b, c): // Infinity
		fallthrough
	case a == 0 && b == 0 && c == 0: // All sides zero
		return NaT

	case a == b && b == c: // Equilateral
		return Equ
	case a == b || b == c || a == c: // Isosceles
		return Iso
	default: // Only thing left would be Scalene
		return Sca
	}
}

func areNaN(vals ...float64) bool {
	for _, v := range vals {
		if math.IsNaN(v) {
			return true
		}
	}

	return false
}

func areInf(vals ...float64) bool {
	for _, v := range vals {
		if math.IsInf(v, 1) || math.IsInf(v, -1) {
			return true
		}
	}

	return false
}
