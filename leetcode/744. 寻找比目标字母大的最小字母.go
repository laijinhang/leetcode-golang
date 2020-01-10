package main

func nextGreatestLetter(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	var res byte
	for i := 0;i < len(letters);i++ {
		if letters[i] > target {
			res = letters[i]
			break
		}
	}
	return res
}