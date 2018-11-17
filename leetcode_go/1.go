package main

import (
	"strings"
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 0;i < n;i++ {
		fmt.Scan(&a[i])
	}
	distance := make([]int, n - 1)
	var a1, b1 int
	for i := 0;i < n - 1;i++ {
		fmt.Scan(&a1, &b1)
		a1--
		b1--
		if a[a1] > a[b1] {
			distance[i] = a[a1] - a[b1]
		} else {
			distance[i] = a[b1] - a[a1]
		}
	}
	sort.Ints(distance)
	fmt.Println(distance[len(distance)-k])
	return
	//var n int
	var s string
	fmt.Scan(&n, &s)
	i1 := strings.Index(s, "msc")
	i2 := strings.Index(s, "mcc")
	if i1 == -1 || i2 == -1 {
		fmt.Println("0")
		return
	}
	mscs := []int{i1}
	mccs := []int{i2}
	for mscs[len(mscs)-1] + 3 < len(s) {
		i1 = strings.Index(s[mscs[len(mscs)-1]+3:], "msc")
		if i1 == -1 {
			break
		}
		i1 += mscs[len(mscs)-1] + 3
		mscs = append(mscs, (i1))
	}
	for mccs[len(mccs)-1] + 3 < len(s) {
		i1 = strings.Index(s[mccs[len(mccs)-1]+3:], "mcc")
		if i1 == -1 {
			break
		}
		i1 += mccs[len(mccs)-1] + 3
		mccs = append(mccs, (i1))
	}
	num := 0
	for i := 0;i < len(mscs);i++ {
		for j := 0;j < len(mccs);j++ {
			if mscs[i] < mccs[j] {
				num++
			}
		}
	}
	fmt.Println(num)
}
