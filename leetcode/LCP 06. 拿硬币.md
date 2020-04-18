[LCP 06. 拿硬币](https://leetcode-cn.com/problems/na-ying-bi/)

```golang
func minCount(coins []int) int {
    cnt := 0
    for i := 0;i < len(coins);i++ {
        cnt += (coins[i] + 1) / 2
    }
    return cnt
}
```
