package is_graph_bipartite

import (
	"fmt"
)

const (
	WHITE = 0 //白色

	GRAY = 1 //灰色

	BLACK = 2 //黑色

	RED = 3 //红色

	GREEN = 4 //绿色
)

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))
	valid := true
	for i := 0; i < len(graph) && valid; i++ {
		if color[i] == WHITE {
			color[i] = RED
			valid = dfsBipartite(i, color, GREEN, graph)
		} else {
			cn := RED
			if color[i] == RED {
				cn = GREEN
			}
			valid = dfsBipartite(i, color, cn, graph)
		}
	}
	return valid
}

func bfsBipartite(s int, color []int, graph [][]int) bool {
	queue := make([]int, 1)
	queue[0] = s
	color[s] = RED
	for ; len(queue) > 0; {
		u := queue[0]
		queue = queue[1:]
		cu := color[u]
		for _, v := range graph[u] {
			cv := color[v]
			if cv == WHITE {
				vInQueue := false
				for _, q := range queue {
					if v == q {
						vInQueue = true
						break
					}
				}
				if !vInQueue {
					queue = append(queue, v)
				}
				if cu == RED {
					color[v] = GREEN
				} else if cu == GREEN {
					color[v] = RED
				} else {
					return false
				}
			} else if cv == RED {
				if cu != GREEN {
					return false
				}
			} else if cv == GREEN {
				if cu != RED {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func dfsBipartite(u int, color []int, pcv int, graph [][]int) bool {
	for _, v := range graph[u] {
		if color[v] == WHITE {
			color[v] = pcv
			npcv := WHITE
			if pcv == RED {
				npcv = GREEN
			} else {
				npcv = RED
			}
			valid := dfsBipartite(v, color, npcv, graph)
			if valid == false {
				return false
			}
		} else if color[v] != pcv {
			return false
		}
	}
	return true
}

//图的广度优先搜索
func bfs(graph [][]int, s int) (color []int, d []int, p []int) {
	color = make([]int, len(graph))
	d = make([]int, len(graph))
	p = make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		color[i] = WHITE
		d[i] = -1
		p[i] = -1
	}
	//初始化第一个节点
	color[s] = GRAY
	d[s] = 0
	p[s] = -1
	queue := make([]int, 1)
	queue[0] = s

	for ; len(queue) > 0; {
		fmt.Println(queue)
		u := queue[0]
		queue = queue[1:]
		for _, v := range graph[u] {
			if color[v] != WHITE {
				continue
			}
			vInQueue := false
			for _, q := range queue {
				if v == q {
					vInQueue = true
					break
				}
			}
			if !vInQueue {
				queue = append(queue, v)
			}
			color[v] = GRAY
			d[v] = d[u] + 1
			p[v] = u
		}
		color[u] = BLACK
	}
	return
}

type Graph struct {
	G     map[string][]string
	D     map[string]int
	F     map[string]int
	P     map[string]string
	Color map[string]int
	time  int
}

func (graph *Graph) Init(g map[string][]string) {
	graph.G = g
	graph.Color = make(map[string]int, len(g))
	graph.D = make(map[string]int, len(g))
	graph.F = make(map[string]int, len(g))
	graph.P = make(map[string]string, len(g))
	graph.time = 0

	for u := range g {
		graph.Color[u] = WHITE
		graph.D[u] = -1
		graph.F[u] = -1
		graph.P[u] = "0"
	}
}

//图的深度优先搜索
func (graph *Graph) dfs() {
	if graph.Color["s"] == WHITE {
		graph.dfsVisit("s")
	}

	for u := range graph.G {
		if graph.Color[u] == WHITE {
			graph.dfsVisit(u)
		}
	}
}

func (graph *Graph) dfsVisit(u string) {
	graph.time = graph.time + 1
	graph.D[u] = graph.time
	graph.Color[u] = GRAY
	for _, v := range graph.G[u] {
		if graph.Color[v] == WHITE {
			graph.P[v] = u
			graph.dfsVisit(v)
		}
	}
	graph.Color[u] = BLACK
	graph.time = graph.time + 1
	graph.F[u] = graph.time
}
