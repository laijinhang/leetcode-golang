/*
题目描述
对于一个元素各不相同且按升序排列的有序序列，请编写一个算法，创建一棵高度最小的二叉查找树。

给定一个有序序列int[] vals,请返回创建的二叉查找树的高度。

 */
package main

import "fmt"

type tree struct {
	val int
	left *tree
	right *tree
}

func (tree) buildMinimalBST(vals []int) int {
	root := tree{}
	bst(&root, vals, 0, len(vals) - 1)

	len1 := len(vals)
	h := 0
	for ;len1 != 0;h++ {
		len1 /= 2
	}
	return h
}

func bst(t *tree, vals []int, l, r int) {
	if l == r {
		return
	}
	// 左边
	r = (l + r) / 2
	t.left = &tree{val:vals[r]}
	bst(t.left, vals, l, r)
	// 右边
	if (l + r) % 2 == 0 {
		l = (l + r) / 2
	} else {
		l = (l + r + 1) / 2
	}
	t.right = &tree{val:vals[l]}
	bst(t.right, vals, l, r)
}

func main() {
	t := tree{}
	h := t.buildMinimalBST([]int{1,2,3,4,5})
	fmt.Println(h)
}