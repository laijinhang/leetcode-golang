package main

import "fmt"

func detectCapitalUse(word string) bool {
	for i := 0;i < len(word);i++ {}
	wLen := len(word)
	if word[0] >= 'a' && word[0] <= 'z' {
		for i := 1;i < wLen;i++ {
			if word[i] >= 'A' && word[i] <= 'Z' {
				return false
			}
		}
	} else if wLen != 1 {
		if word[1] >= 'A' && word[1] <= 'Z' {
			for i := 2;i < wLen;i++ {
				if word[i] >= 'a' && word[i] <= 'z' {
					return false
				}
			}
		} else {
			for i := 2;i < wLen;i++ {
				if word[i] >= 'A' && word[i] <= 'Z' {
					return false
				}
			}
		}
	}
	return true
}

func main()  {
	s := []string{"Wda", "sdas", "ADSD", "dasA", "AdsA"}
	for _, str := range s {
		fmt.Println(detectCapitalUse(str))
	}
}