package main

import "fmt"

func lemonadeChange(bills []int) bool {
	len1 := len(bills)
	money := [3]int{0, 0, 0}
	for i := 0;i < len1;i++ {
		switch bills[i] {
		case 5:
			money[0]++
		case 10:
			if money[0] == 0 {
				return false
			}
			money[0]--
			money[1]++
		case 20:
			if money[0] == 0 || (money[0] < 3 && money[1] == 0) {
				return false
			}
			if money[1] == 0 {
				money[0] -= 3
			} else {
				money[0]--
				money[1]--
			}
			money[2]++
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5,5,10,20,5,5,5,5,5,5,5,5,5,10,5,5,20,5,20,5}))
}