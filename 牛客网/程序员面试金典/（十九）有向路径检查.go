/*
题目描述
对于一个有向图，请实现一个算法，找出两点之间是否存在一条路径。

给定图中的两个结点的指针DirectedGraphNode* a, DirectedGraphNode* b(请不要在意数据类型，图是有向图),请返回一个bool，代表两点之间是否存在一条路径(a到b或b到a)。

分析：
从题目中可以看出，其实就是图的遍历，有两种遍历方式，广度优先和深度优先

这边我是使用广度优先来解的：
代码实现：
 */
package main

type UndirectedGraphNode struct {
	val interface{}
	neighbors []*UndirectedGraphNode
}

func checkPath(a, b *UndirectedGraphNode) bool {
	var goThrough map[*UndirectedGraphNode]bool	// 走过的路径
	nodeQueue := make([]*UndirectedGraphNode, 0)

	nodeQueue = append(nodeQueue, a)
	for len(nodeQueue) == 0 {
		temp := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		for _, node := range temp.neighbors {
			if node == b {
				return true
			}
			if goThrough[node] != true {	// 没有走过
				nodeQueue = append(nodeQueue, node)
			}
		}
		goThrough[temp] = true	// 当前节点走过
	}

	return false
}

func main() {

}