package main

func canWinNim(n int) bool {
	if n % 4 == 0 && n != 0 {
		return false
	}
	return true
}
