package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	resTime := 0
	for i := 0;i < len(timeSeries) - 1;i++ {
		if timeSeries[i+1] - timeSeries[i] < duration {
			resTime += timeSeries[i+1] - timeSeries[i]
		} else {
			resTime += duration
		}
	}
	if len(timeSeries) > 0 {
		resTime += duration
	}
	return resTime
}