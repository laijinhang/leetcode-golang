package main

func countAndSay(n int) string {
	bs := []byte{}
	for i := 1;i < n;i++ {
		temp := []byte{}
		for j, k := 0, 0;j < len(bs);j = k {
			for k = j + 1;k < len(bs) && bs[j] == bs[k];k++ {}
			temp = append(temp, byte(k - j + '0'), bs[j])
		}
		bs = temp
	}
	return string(bs)
}

func main()  {
	countAndSay(3)
}
