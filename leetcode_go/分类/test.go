package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0;i < 10;i++ {
		fmt.Println(rand.Intn(2)) // => 134020434
	}
}