package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	if N % 2 == 0 { return nil}
	if N == 1 {return []*TreeNode{&TreeNode{}}}
	ans := []*TreeNode{}
	for i := 1;i < N;i += 2 {
		for _, l := range allPossibleFBT(i) {
			for _, r := range allPossibleFBT(N - i - 1) {
				var root TreeNode
				root.Left = l
				root.Right = r
				ans = append(ans, &root)
			}
		}
	}
	return ans
}

func main() {
}

