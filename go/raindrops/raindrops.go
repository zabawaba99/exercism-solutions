// +build !example

package raindrops

import "strconv"

const TestVersion = 1

func Convert(drop int) string {
	var s string
	if drop%3 == 0 {
		s += "Pling"
	}

	if drop%5 == 0 {
		s += "Plang"
	}

	if drop%7 == 0 {
		s += "Plong"
	}

	if s == "" {
		s = strconv.FormatInt(int64(drop), 10)
	}
	return s
}
