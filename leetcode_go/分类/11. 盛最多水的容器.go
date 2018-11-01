package main

import "math"

func maxArea(height []int) int {
	maxArea, area := 0, 0
	for i := 0;i < len(height);i++ {
		for j := i + 1;j < len(height);j++ {
			area = (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}