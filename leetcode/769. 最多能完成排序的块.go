package main

func maxChunksToSorted(arr []int) int {
	min, max, areaMax := 0, 0, 0
	num := 0
	for i, j := 0, 0;i < len(arr);i = min + 1 {
		min = i
		max = i
		areaMax = i
		for j = i + 1;j < len(arr);j++ {
			if arr[j] < arr[areaMax] || (arr[j] == arr[areaMax] && areaMax - j != 1) {
				areaMax = max
				min = j
			}
			if arr[j] > arr[max]  {
				max = j
			}
		}
		num++
	}
	return num
}