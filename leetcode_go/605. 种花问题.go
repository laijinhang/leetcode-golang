package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	// 长度为1
	count := 0
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			count++
		}
	} else {
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			flowerbed[0] = 1
			count++
		}
		for i := 2; i < len(flowerbed) - 1; i++ {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count++
			}
		}
		if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
			flowerbed[len(flowerbed)-1] = 1
			count++
		}
	}
	if count < n {
		return false
	}
	return true
}