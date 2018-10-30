package main

func stoneGame(piles []int) bool {
	a, b := 0, 0
	bi, ei := 0, len(piles)
	i := 0
	for ;a < b; {
		if piles[bi] >= piles[ei] {
			if i % 2 == 0 {
				a += piles[bi]
			} else {
				b += piles[bi]
			}
			bi++
		} else {
			if i % 2 == 0 {
				a += piles[ei]
			} else {
				b += piles[ei]
			}
			ei++
		}
		i++
	}
	if b > a {
		return false
	}
	return true
}