// Go strings are read-only slices of bytes.  The language treats them as containers of UTF-8 encoded text.
// In Go, characters are considered 'runes', an integer that represents a Unicode code point.
// https://go.dev/blog/strings
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Strings are UTF-8 encoded text
	const s = "สวัสดีt"

	// As they are equivalent to []byte, len() of strings will produce length of raw bytes stored within
	fmt.Println("len:", len(s))

	// Indexing into a string produces raw byte values at each index
	// This generates raw hex values of bytes that constitute the code in constant "s"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count runes, use the utf8 package.
	// NOTE: run-time depends on size of string as RuneCountInString has to decode each rune
	// Some language characters are represented by UTF-8 code points that span multiple bytes, so count may be higher than expected!
	fmt.Println("rune count:", utf8.RuneCountInString(s))

	// Range loop handles strings specially and decodes each run along with offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
		examineRune(runeValue)
	}

	// Can achieve the same as above by using the utf8.DecodeRuneInString
	fmt.Println("\nUsing utf8.DecodeRuneInString()")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}

// Values in quotes in
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	} else {
		fmt.Println("is neither tee or so sua")
	}
}
