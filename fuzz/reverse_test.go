package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, err := Reverse(tc.in)
		if err != nil {
			t.Errorf("invalid utf-8 character")
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello", "world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig)
		if err != nil {
			t.Skip("skiping this test vecause invalid utf-8 character", orig)
		}
		doubleRev, err := Reverse(rev)
		if err != nil {
			t.Skip("skiping this test vecause invalid utf-8 character", orig)
		}
		t.Logf("Number of rues: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced an invalid UTF-8 string %q", rev)
		}
	})
}

func FuzzDivide(f *testing.F) {

	for i := -10; i < 10; i++ {
		f.Add(i, 1+100)

	}

	f.Fuzz(func(t *testing.T, a int, b int) {
		result := Divide(a, b)
		t.Log("the result is", result)

	})
}
