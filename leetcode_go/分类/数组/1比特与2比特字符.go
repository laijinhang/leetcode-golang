package main

func isOneBitCharacter(bits []int) bool {
	res := true
	for i := 0;i < len(bits);i++ {
		if bits[i] == 1 {
			i++
			if i == len(bits) - 1 {
				res = false
			} else if i == len(bits) - 2 {
				res = true
			}
		}
	}
	return res
}