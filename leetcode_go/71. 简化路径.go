package main

import (
	"fmt"
	"path"
)

func main() {
	var s string
	for {
		fmt.Scan(&s)
		res := path.Dir(s)
		fmt.Println(res)
		_, fileName := path.Split(s)
		if fileName != "." || fileName != ".." {
			res = path.Join(res, fileName)
		}
		if res == "." {
			res = path.Join("/")
		}
		fmt.Println(res)
	}
}
