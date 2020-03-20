[最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

```golang
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
```
