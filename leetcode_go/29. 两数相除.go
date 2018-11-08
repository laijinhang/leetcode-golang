package main

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	res := 0
	if dividend >= 0 && divisor > 0 {
		for ;divisor <= dividend;dividend -= divisor {
			res++
		}
	} else {
		if (dividend <= 0 && divisor > 0) {
			dividend = -dividend
			for ;divisor <= dividend;dividend -= divisor {
				res++
			}
			res = -res
		} else if (dividend >= 0 && divisor < 0) {
			divisor = -divisor
			for ;divisor <= dividend;dividend -= divisor {
				res++
			}
			res = -res
		} else {
			divisor = -divisor
			dividend = -dividend
			for ;divisor <= dividend;dividend -= divisor {
				res++
			}
		}
	}
	return res
}
