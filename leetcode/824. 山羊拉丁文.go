package main

import (
	"strings"
	"fmt"
)

var vowel = map[byte]bool {
	'a': true,
	'A': true,
	'e': true,
	'E': true,
	'i': true,
	'I': true,
	'o': true,
	'O': true,
	'u': true,
	'U': true}

func toGoatLatin(S string) string {
	temps := strings.Split(S, " ")
	a := ""
	for key, s := range temps {
		a += "a"
		temps[key] = func(b []byte) string {
			if !vowel[b[0]] {
				b = append(b[1:], b[0])
			}
			b = append(b, 'm', 'a')
			return string(b)
		}([]byte(s)) + a
	}
	return strings.Join(temps, " ")
}

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}