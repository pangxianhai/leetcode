package is_graph_bipartite

import (
	"fmt"
	"testing"
)

func TestIsBipartite(t *testing.T) {
	var graph [][]int
	//graph = make([][]int, 5)
	//graph[0] = []int{1, 4}
	//graph[1] = []int{0, 2, 3, 4}
	//graph[2] = []int{1, 3}
	//graph[3] = []int{1, 2, 4}
	//graph[4] = []int{0, 1, 3}
	//fmt.Println(false, isBipartite(graph))
	//
	//graph = make([][]int, 10)
	//graph[0] = []int{}
	//graph[1] = []int{2, 4, 6}
	//graph[2] = []int{1, 4, 8, 9}
	//graph[3] = []int{7, 8}
	//graph[4] = []int{1, 2, 8, 9}
	//graph[5] = []int{6, 9}
	//graph[6] = []int{1, 5, 7, 8, 9}
	//graph[7] = []int{3, 6, 9}
	//graph[8] = []int{2, 3, 4, 6, 9}
	//graph[9] = []int{2, 4, 5, 6, 7, 8}
	//fmt.Println(false, isBipartite(graph))
	//
	//graph = make([][]int, 10)
	//graph[0] = []int{3, 4, 6}
	//graph[1] = []int{3, 6}
	//graph[2] = []int{3, 6}
	//graph[3] = []int{0, 1, 2, 5}
	//graph[4] = []int{0, 7, 8}
	//graph[5] = []int{3}
	//graph[6] = []int{0, 1, 2, 7}
	//graph[7] = []int{4, 6}
	//graph[8] = []int{4}
	//graph[9] = []int{}
	//fmt.Println(true, isBipartite(graph))

	graph = make([][]int, 5)
	graph[0] = []int{4, 1}
	graph[0] = []int{0, 2}
	graph[0] = []int{1, 3}
	graph[0] = []int{2, 4}
	graph[0] = []int{3, 0}
	fmt.Println(false, isBipartite(graph))

}

func TestGraphBfs(t *testing.T) {
	graph := make([][]int, 5)
	graph[0] = []int{1, 4}
	graph[1] = []int{0, 2, 3, 4}
	graph[2] = []int{1, 3}
	graph[3] = []int{1, 2, 4}
	graph[4] = []int{0, 1, 3}
	color, d, p := bfs(graph, 3)

	fmt.Println("颜色", color)
	fmt.Println("路劲", d)
	fmt.Println("父级", p)
}

func TestGraphDfs(t *testing.T) {
	graph := map[string][]string{
		"s": {"z", "w"},
		"z": {"y", "w"},
		"y": {"x"},
		"x": {"z"},
		"w": {"x"},
		"v": {"w", "s"},
		"t": {"v", "u"},
		"u": {"t", "v"},
	}
	g := Graph{}
	g.Init(graph)
	g.dfs()

	fmt.Println("d:", g.D)
	fmt.Println("F:", g.F)
	fmt.Println("p:", g.P)
}
