// Package textstats counts characters, words, vowels and other
// categories in a piece of text. It is Unicode-aware: a character
// is counted as a rune, not as a byte.
package textstats

import (
	"strings"
	"unicode"
)

// Stats is the result of analysing a text.
type Stats struct {
	Bytes       int // raw UTF-8 bytes (what len() returns)
	Characters  int // all runes, including spaces
	NonSpace    int // runes that are not whitespace
	Words       int
	Vowels      int // English vowels a e i o u (either case)
	Consonants  int // English letters that are not vowels
	Digits      int
	Spaces      int
	Punctuation int
	Lines       int
}

// Analyze counts every category in text in a single pass over its runes.
func Analyze(text string) Stats {
	s := Stats{Bytes: len(text), Words: len(strings.Fields(text))}
	if text != "" {
		s.Lines = strings.Count(text, "\n") + 1
	}
	for _, r := range text { // range over a string yields runes
		s.Characters++
		switch {
		case unicode.IsSpace(r):
			s.Spaces++
			continue
		case IsVowel(r):
			s.Vowels++
		case r < unicode.MaxASCII && unicode.IsLetter(r):
			s.Consonants++
		case unicode.IsDigit(r):
			s.Digits++
		case unicode.IsPunct(r):
			s.Punctuation++
		}
		s.NonSpace++
	}
	return s
}

// IsVowel reports whether r is an English vowel (a, e, i, o, u).
func IsVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}
