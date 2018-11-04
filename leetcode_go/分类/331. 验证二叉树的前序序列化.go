package main

import (
    "fmt"
    "strings"
)

func isValidSerialization(preorder string) bool {
    ps := strings.Split(preorder, ",")
    if len(ps) == 0 || len(ps) % 2 != 1 {
        return false
    }
    stack := []int{}
    if ps[0] != "#" {
        stack = append(stack, 1)
    }
    i := 0
    for ;i + 2 < len(ps);i += 2 {
        if len(stack) == 0 {
            return false
        }
        stack = stack[1:]
        if ps[i+1] != "#" {
            stack = append(stack, 1)
        }
        if ps[i+2] != "#" {
            stack = append(stack, 1)
        }
    }
    if len(stack) != 0 {
        return false
    }
    return true
}

func main() {
    fmt.Println(isValidSerialization("1,#,#"))
}