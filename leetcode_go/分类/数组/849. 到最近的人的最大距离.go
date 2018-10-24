package main

func maxDistToClosest(seats []int) int {
	ys := []int{}
	temp := 0
	for i := 0;i < len(seats);i++ {
		if seats[i] == 0 {
			temp++
		} else {
			if len(ys) == 0 && seats[0] == 0 {
				ys = append(ys, temp)
			} else {
				ys = append(ys, (temp + 1) / 2)
			}
			temp = 0
		}
	}
	if seats[len(seats)-1] == 0 {
		ys = append(ys, temp)
	} else {
		ys = append(ys, (temp + 1) / 2)
	}
	max := ys[0]
	for i := 1;i < len(ys);i++ {
		if ys[i] > max {
			max = ys[i]
		}
	}
	return max
}