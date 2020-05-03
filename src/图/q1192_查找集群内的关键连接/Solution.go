/**
 * 强连接下的tarjan算法
 * 中心思想是找出图中的环
 * 首先像树一样深度搜索，记录访问的时间戳来代表树的深度
 * 如果发现子节点访问过，则说明连接成环，于是将树的深度统一，并通过递归向上传递
 * 最后，时间戳不同的节点间即为关键路径
 *
 * 此题中把边的关系转化成类似二维数组，记录每个点可以到达的点，类似矩阵
 */
package main

// Clock 全局变量，每访问一个节点Clock++
var Clock int
// dfn 每个节点的访问时间戳
var dfn []int
// low 每个节点的最小时间戳
var low []int
// visited 记录每个节点是否被访问
var visited []bool
// result 结果
var result [][]int
// nodeSet 将边的关系转化为每个点对应能访问到的节点的关系
var nodeSet [][]int

func criticalConnections(n int, connections [][]int) [][]int {
	// 初始化
	Clock = 0
	dfn = make([]int, n)
	low = make([]int, n)
	visited = make([]bool, n)

	// 二维数组：外层为所有节点，内层为该节点连接的节点
	nodeSet = make([][]int, n)
	for _, edge := range connections {
		nodeSet[edge[0]] = append(nodeSet[edge[0]], edge[1])
		nodeSet[edge[1]] = append(nodeSet[edge[1]], edge[0])
	}

	// 从0节点开始深度搜索
	tarjan(0, -1)
	return result
}

func tarjan(node int, parent int) {
	// 每访问一个节点，Clock++
	// 并且初始化访问时间戳和最小时间戳为Clock
	Clock++
	dfn[node] = Clock
	low[node] = Clock
	visited[node] = true

	// 访问该节点的所有子节点
	for _, child := range nodeSet[node] {
		// 如果该节点为父节点，则忽略
		if child == parent {
			continue
		}

		// 如果没被访问过，则进行访问
		if visited[child] == false {
			tarjan(child, node)

			// 如果该节点的访问时间戳大于子节点，说明连接成环，同步时间戳
			// 否则说明该点到下一个点没有成环，出现了关键路径
			if low[node] >= low[child] {
				low[node] = low[child]
			} else {
				// 
				result = append(result, []int{node, child})
			}

		} else {
			// 如果访问过，说明连接成环
			low[node] = low[child]
		}
	}

}

func main() {
	ans := criticalConnections(6, [][]int{{0,1}, {1,2}, {2,0}, {1,3}, {3,4}, {4,5}, {5,3}})
	for _, i := range ans{
		println("(", i[0], ",", i[0], ")")
	}
}