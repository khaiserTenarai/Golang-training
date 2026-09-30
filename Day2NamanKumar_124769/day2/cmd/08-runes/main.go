// Task 8: Bytes vs runes, using Indian and Unicode names.
//
// A Go string is a sequence of BYTES encoded as UTF-8.
// A rune is one Unicode CODE POINT (an int32).
// English letters take 1 byte; Indian scripts take 3 bytes per code point.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	names := []struct{ label, text string }{
		{"English", "Priya"},
		{"Kannada", "ಕನ್ನಡ"},
		{"Hindi", "नमस्ते"},
		{"Tamil", "தமிழ்"},
		{"Telugu", "తెలుగు"},
		{"Bengali", "বাংলা"},
		{"Accented", "José"},
		{"Emoji", "Go🚀"},
	}

	fmt.Println("=== len() counts BYTES, utf8.RuneCountInString counts RUNES ===")
	fmt.Printf("%-9s %-10s %6s %6s\n", "Script", "Text", "Bytes", "Runes")
	for _, n := range names {
		fmt.Printf("%-9s %-10s %6d %6d\n", n.label, n.text, len(n.text), utf8.RuneCountInString(n.text))
	}

	name := "ಕನ್ನಡ"
	fmt.Printf("\n=== Looping over %q byte by byte (s[i]) ===\n", name)
	for i := 0; i < len(name); i++ {
		fmt.Printf("%02X ", name[i])
	}
	fmt.Println("\n-> 15 raw bytes; printing them as characters gives garbage")

	fmt.Printf("\n=== Looping over %q rune by rune (for range) ===\n", name)
	for i, r := range name {
		fmt.Printf("byte index %2d  rune %-3c  U+%04X  (%d bytes)\n", i, r, r, utf8.RuneLen(r))
	}
	fmt.Println("-> index jumps by 3 because each Kannada code point is 3 bytes")

	fmt.Println("\n=== Converting between string, []byte and []rune ===")
	bs := []byte(name)
	rs := []rune(name)
	fmt.Println("[]byte:", bs)
	fmt.Println("[]rune:", rs)
	fmt.Println("len([]byte) =", len(bs), " len([]rune) =", len(rs))

	fmt.Println("\n=== The slicing bug ===")
	hindi := "नमस्ते"
	fmt.Printf("hindi[:2]          = %q  (broken: cuts a character in half)\n", hindi[:2])
	fmt.Printf("valid UTF-8?         %v\n", utf8.ValidString(hindi[:2]))
	fmt.Printf("string([]rune)[:2] = %q  (correct: first 2 runes)\n", string([]rune(hindi)[:2]))

	fmt.Println("\n=== Reversing: bytes vs runes ===")
	fmt.Printf("reverse by bytes : %q  (invalid)\n", reverseBytes("Priya ಕನ್ನಡ"))
	fmt.Printf("reverse by runes : %q\n", reverseRunes("Priya ಕನ್ನಡ"))

	fmt.Println("\n=== Runes are not always visible letters ===")
	fmt.Println(`"ಕನ್ನಡ" looks like 3 letters (ಕ ನ್ನ ಡ) but has 5 runes:`)
	for _, r := range "ಕನ್ನಡ" {
		fmt.Printf("  U+%04X %q\n", r, r)
	}
	fmt.Println("The virama (U+0CCD) joins consonants into one visible letter.")
	fmt.Println("Counting what humans see (grapheme clusters) needs a Unicode segmentation library.")

	fmt.Println("\n=== byte vs rune literals ===")
	var b byte = 'P'
	var r rune = 'ಕ'
	fmt.Printf("byte 'P' = %d (%T)\nrune 'ಕ' = %d (%T)\n", b, b, r, r)
}

func reverseBytes(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func reverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
