package main

import "fmt"

type longScan struct {
	num int	// 次数
	char byte // 字符
}

func isLongPressedName(name string, typed string) bool {
	nScan, tScan := []longScan{}, []longScan{}
	nLen, tLen := len(name), len(typed)
	if nLen == 0 && tLen == 0 {
		return true
	} else if (nLen != 0 && tLen == 0) || nLen > tLen {
		return false
	}
	nScan = append(nScan, longScan{num:1, char:name[0]})
	for i := 1;i < nLen;i++ {
		if name[i] == nScan[len(nScan)-1].char {
			nScan[len(nScan)-1].num++
		} else {
			nScan = append(nScan, longScan{num:1, char:name[i]})
		}
	}
	tScan = append(tScan, longScan{num:1, char:typed[0]})
	for i := 1;i < tLen;i++ {
		if typed[i] == tScan[len(tScan)-1].char {
			tScan[len(tScan)-1].num++
		} else {
			tScan = append(tScan, longScan{num:1, char:typed[i]})
		}
	}
	if len(nScan) != len(tScan) {
		return false
	}
	for i := 0;i < len(nScan);i++ {
		if nScan[i].char != tScan[i].char || nScan[i].num > tScan[i].num {
			return false
		}
	}
	return true
}

func main()  {
	fmt.Println(isLongPressedName("leelee", "lleeelee"))
}