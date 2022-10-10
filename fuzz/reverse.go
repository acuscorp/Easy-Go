package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	doubleRev, _ := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not a valid UTF-8")
	}
	// fmt.Printf("input: %q\n", s)
	r := []rune(s)
	// fmt.Printf("runes: %q\n", r)
	fmt.Println()
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]

	}
	return string(r), nil
}
