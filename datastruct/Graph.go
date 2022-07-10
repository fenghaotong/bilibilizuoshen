package datastruct

type Graph struct {
	Nodes map[int]GraphNode
	Edges []GraphEdge
}

type GraphNode struct {
	Val   Element
	In    int
	Out   int
	Next  *[]GraphNode
	Edges *[]GraphEdge
}

type GraphEdge struct {
	Weight int
	From   GraphNode
	To     GraphNode
}

func (edge *GraphEdge) GraphEdgeEqual(cur GraphEdge) bool {
	if edge.Weight == cur.Weight && edge.From == cur.From && edge.To == cur.To {
		return true
	}
	return false
}

func ContainsKey(m map[int]GraphNode, key int) bool {
	for k, _ := range m {
		if k == key {
			return true
		}
	}
	return false
}

func ContainsList(list []GraphEdge, edge GraphEdge) bool {
	for _, v := range list {
		if v.GraphEdgeEqual(edge) {
			return true
		}
	}
	return false
}

// [[fromNode, toNode, weight]]
func CreateGraph(matrix [][]int) *Graph {
	graph := &Graph{}
	for i := 0; i < len(matrix); i++ {
		from := matrix[i][0]
		to := matrix[i][1]
		weight := matrix[i][2]
		if !ContainsKey(graph.Nodes, from) {
			graph.Nodes[from] = GraphNode{Val: Element(from)}
		}

		if !ContainsKey(graph.Nodes, to) {
			graph.Nodes[to] = GraphNode{Val: Element(to)}
		}

		fromNode := graph.Nodes[from]
		toNode := graph.Nodes[to]
		newEdge := GraphEdge{weight, fromNode, toNode}
		*fromNode.Next = append(*fromNode.Next, toNode)
		fromNode.Out++
		toNode.In++
		*fromNode.Edges = append(*fromNode.Edges, newEdge)
		if !ContainsList(graph.Edges, newEdge) {
			graph.Edges = append(graph.Edges, newEdge)
		}
	}

	return graph
}
