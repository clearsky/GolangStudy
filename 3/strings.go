package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s :=  "Yes我爱慕课网"
	fmt.Printf("%s, %d", s, len(s))
	fmt.Println()
	for i, ch := range s{  // ch is a rune, unicode
		fmt.Printf("(%d %x)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune Count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0{
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s){
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
