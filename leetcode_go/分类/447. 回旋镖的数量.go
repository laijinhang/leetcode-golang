package main

func numberOfBoomerangs(points [][]int) int {
	num := 0
	for i := 0;i < len(points);i++ {
		for j := i + 1;j < len(points);j++ {
			for k := j + 1;k < len(points);k++ {
				if (points[i][0] - points[j][0]) * (points[i][0] - points[j][0]) +
					(points[i][1] - points[j][1]) * (points[i][1] - points[j][1]) ==
					(points[i][0] - points[k][0]) * (points[i][0] - points[k][0]) +
						(points[i][1] - points[k][1]) * (points[i][1] - points[k][1]) {
							num++
				} else if (points[j][0] - points[i][0]) * (points[j][0] - points[i][0]) +
					(points[j][1] - points[i][1]) * (points[j][1] - points[i][1]) ==
					(points[j][0] - points[k][0]) * (points[j][0] - points[k][0]) +
						(points[j][1] - points[k][1]) * (points[j][1] - points[k][1]) {
							num++
				} else if (points[k][0] - points[i][0]) * (points[k][0] - points[i][0]) +
					(points[k][1] - points[i][1]) * (points[k][1] - points[i][1]) ==
					(points[k][0] - points[j][0]) * (points[k][0] - points[j][0]) +
						(points[k][1] - points[j][1]) * (points[k][1] - points[j][1]) {
							num++
				}
			}
		}
	}
	return num * 2
}