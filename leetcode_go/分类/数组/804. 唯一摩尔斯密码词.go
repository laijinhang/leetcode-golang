package main

import "fmt"

var pw = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]bool)
	for _, word := range words {
		temp := ""
		for i := 0;i < len(word);i++ {
			temp += pw[word[i]-'a']
		}
		m[temp] = true
	}
	return len(m)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}