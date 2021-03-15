[233. 数字 1 的个数](https://leetcode-cn.com/problems/number-of-digit-one/)
```golang
func countDigitOne(n int) int {
	res, base := 0, 1
	height, cur, low := n, n, n
	for base <= n {
		low = n % base
		height = n / base
		cur = height % 10
		height /= 10
		if cur == 0 {
			res += height * base
		} else if cur == 1 {
			res += height*base + low + 1
		} else {
			res += (height + 1) * base
		}
		base *= 10
	}
	return res
}
```