package main

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	tempStrB := make([][]byte, numRows)
	index := 0
	direction := 1
	for i := 0;i < len(s);i++ {
		tempStrB[index] = append(tempStrB[index], s[i])
		index += direction
		if index == numRows - 1 {
			direction = -1
		} else if index == 0 {
			direction = 1
		}
	}
	str := ""
	for i := 0;i < len(tempStrB);i++ {
		str += string(tempStrB[i])
	}
	return str
}