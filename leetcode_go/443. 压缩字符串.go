package main

func compress(chars []byte) int {
	charNum := 0	// 存放字符个数
	moveLen := 0
	for i := 0;i < len(chars) - moveLen;i++ {
		j := i + 1
		for ;j < len(chars) - moveLen && chars[i] == chars[j];j++ {}
		if j == i + 1 {
			continue
		} else {
			// 向前移动
			charNum = j - i
			k := j - 1
			for cn := charNum;k > i && cn != 0;k, cn = k - 1, cn / 10 {
				chars[k] = byte(cn % 10) + '0'
			}
			// 向前移动
			m := 1
			for ;k + m < len(chars);m++ {
				chars[i+m] = byte(chars[k+m])
			}
			moveLen += k - i
			if k != i && j != len(chars) {
				for ;charNum != 0;charNum /= 10 {
					i++
				}
			} else {
				i = j - 1
			}
		}
	}
	return len(chars) - moveLen
}

func main() {
	compress([]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'})
}