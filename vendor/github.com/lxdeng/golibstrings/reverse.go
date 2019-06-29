// Package contains utility functions for working with strings.
package golibstrings

import (
	"fmt"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	fmt.Println("Enter Reverse")

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
