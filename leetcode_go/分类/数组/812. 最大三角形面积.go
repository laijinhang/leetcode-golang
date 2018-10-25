package main

import "math"

func largestTriangleArea(points [][]int) float64 {
	res := 0.0
	for i := 0;i < len(points);i++ {
		for j := i + 1;j < len(points);j++ {
			for k := j + 1;k < len(points);k++ {
				res = math.Max(res, 0.5 *
					math.Abs(float64(points[i][0] * (points[j][1] - points[k][1]) + points[j][0] * (points[k][1] - points[i][1]) + points[k][0] * (points[i][1] - points[j][1]))))
			}
		}
	}
	return res
}