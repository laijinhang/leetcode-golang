package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	var b1, b2, b3 bytes.Buffer
	tb := time.Now()
	for i := 0;i < 1000;i++ {
		b1.WriteString("abc")
		b2.WriteString("abe")
		b3.WriteString("abf")
		if b1.String() == b2.String() && b2.String() == b3.String() {
			break
		}
	}
	fmt.Println("Run: ", time.Since(tb))

	tb = time.Now()
	var s1, s2, s3 string
	for i := 0;i < 1000;i++ {
		s1 = "abc"
		s2 = "abe"
		s3 = "abf"
		if s1 == s2 && s2 == s3 {
			break
		}
	}
	fmt.Println("Run: ", time.Since(tb))
}