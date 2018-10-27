package main

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	flag := false
	getFactor(num, &flag)
	return flag
}

func getFactor(num int, flag *bool) {
	if num == 2 || num == 3 || num == 5 || *flag {
		*flag = true
		return
	}
	if num % 2 == 0 {
		getFactor(num / 2, flag)
	}
	if num % 3 == 0 {
		getFactor(num / 3, flag)
	}
	if num % 5 == 0 {
		getFactor(num / 5, flag)
	}
}

func main() {
	createTablePrime(100)
}