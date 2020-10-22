// Copyright (c) 2019-2020 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

// Package ascii -- provides ASCII-only character functions.
package ascii

// Codepoint -- returns true iff `c` is an ASCII character.
func Codepoint(c byte) bool {
	return c <= 0x7F
}

// Control -- returns true iff `c` is in a control character.
func Control(c byte) bool {
	return c <= 0x1F || c == 0x7F
}

// Alphabetic -- returns true iff `c` is a letter character.
func Alphabetic(c byte) bool {
	return (0x41 <= c && c <= 0x5A) || (0x61 <= c && c <= 0x7A)
}

// Numeric -- returns true iff `c` is a digit character.
func Numeric(c byte) bool {
	return 0x30 <= c && c <= 0x39
}

// Alphanumeric -- returns true iff `c` is a letter or a digit character.
func Alphanumeric(c byte) bool {
	return (0x41 <= c && c <= 0x5A) || (0x61 <= c && c <= 0x7A) || (0x30 <= c && c <= 0x39)
}

// Uppercase -- returns true iff `c` is an upper case letter character.
func Uppercase(c byte) bool {
	return 0x41 <= c && c <= 0x5A
}

// Lowercase -- returns true iff `c` is a lower case letter character.
func Lowercase(c byte) bool {
	return 0x61 <= c && c <= 0x7A
}

// Upcase -- returns `c` converted to its upper case equivalent;
// when `c` is not a lower case letter character, returns `c` unchanged.
func Upcase(c byte) byte {
	if Lowercase(c) {
		return c - 0x20
	}
	return c
}

// Downcase -- returns `c` converted to its lower case equivalent;
// when `c` is not an upper case letter character, returns `c` unchanged.
func Downcase(c byte) byte {
	if Uppercase(c) {
		return c + 0x20
	}
	return c
}

// CompareCI -- performs case-insensitive character comparison. Upper case letter characters
// are converted to theier lower case equivalents and then compared numerically. Returns one
// of {-1,0,+1} whether correspondingly `c1` is less than, equal to, or greater than `c2`.
func CompareCI(c1, c2 byte) int {
	c1 = Downcase(c1)
	c2 = Downcase(c2)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return +1
	default:
		return 0
	}
}
