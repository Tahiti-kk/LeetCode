// 强连接下的tarjan算法
package main

// Clock 全局变量，每访问一个节点Clock++
var Clock int
// 访问时间戳
var dfn []int
// 最小时间戳
var low []int
// 查看每个节点是否被访问
var visited []bool
var result [][2]int

var nodeSet [][]int

func criticalConnections(n int, connections [][]int) [][]int {
	Clock = 0
	dfn = make([]int, n)
	low = make([]int, n)
	visited = make([]bool, n)

	// 二维数组：外层为所有节点，内层为该节点连接的节点
	nodeSet = make([][]int, n)
	for i:= 0; i<len(nodeSet); i++ {
		nodeSet[i] = make([]int, n)
	}
	for _, edge := range connections {
		nodeSet[edge[0]] = append(nodeSet[edge[0]], edge[1])
		nodeSet[edge[1]] = append(nodeSet[edge[1]], edge[0])
	}

	// 遍历每个节点的可连接点数组
	tarjan(0, -1)
	return result
}

func tarjan(node int, pnode int) {
	Clock++
	dfn[node] = Clock
	low[node] = Clock
	visited[node] = true
	// 访问该节点的所有子节点
	for cnode := range nodeSet[node] {
		// 忽略
		if cnode == pnode {
			continue
		}

		// 如果没被访问过，则进行访问
		if visited[cnode] == false {
			tarjan(cnode, node)

			low[node] = min(low[node], low[cnode])
			// 如果
			if low[node] < low[cnode] {
				result = append(result, []int{node, cnode})
			}

		} else {
			low[node] = min(low[node], low[cnode])
		}
	}

}

func min(x int, y int) int {
	if x<y {
        return x
    }
    return y
}