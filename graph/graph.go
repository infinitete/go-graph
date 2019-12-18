package graph

const BitsPerWord = 32 << (^uint(0) >> 63) // either 32 or 64

const (
	MaxInt  = 1<<(BitsPerWord-1) - 1 // either (1<<31)-1 or (1<<63)-1
	MinInt  = -MaxInt - 1            // either -1<<31 or -1<<63
	MaxUint = 1<<BitsPerWord - 1     // either (1<<32)-1 or (1<<64)-1
)

// 点
type Point struct {

	// 内容
	Name string

	// 边
	Edges []*Edge
}

func (p *Point) AddEdge(point *Point, distance int) {

	for _, e := range p.Edges {
		if e.Point == point {
			return
		}
	}

	edge := &Edge{
		Point:    point,
		Distance: distance,
	}

	p.Edges = append(p.Edges, edge)
}

type Edge struct {
	Point    *Point
	Distance int
}

type Graph struct {
	Points           map[*Point]struct{}
	MinimalDistances map[*Point]map[*Point]int // [from][to]distance
}

func (g *Graph) HasPoint(p *Point) bool {
	_, ok := g.Points[p]
	return ok
}

// 构建地一层直接关联边的路径距离数据
func (g *Graph) BuildMinimalDistance() map[*Point]map[*Point]int {

	distances := make(map[*Point]map[*Point]int)

	for point, _ := range g.Points {
		for _, edge := range point.Edges {
			if distances[point] == nil {
				distances[point] = make(map[*Point]int)
			}
			distances[point][edge.Point] = edge.Distance
		}
	}

	g.MinimalDistances = distances

	return distances
}

// 两点最短路径
func (g *Graph) MinDistance(from, dest *Point, distance int) int {

	// 点不在图中
	if !g.HasPoint(from) || !g.HasPoint(dest) {
		return -1
	}

	if from == dest {
		return distance
	}

	for p, d := range g.MinimalDistances[from] {

		// 如果有直接关联，则距离设置成边的距离
		if p == dest {
			distance = d
		}

		// 如果点p没有其他向下关联的边，则不再向下搜索
		// 此时返回distance即可
		// 贪婪模式终止
		if _, ok := g.MinimalDistances[p]; !ok {
			continue
		}

		// 继续向下搜索，能够找到更短路径
		// 则找每次寻找最短距离，直到找到为止
		// 继续贪婪模式
		if p != dest {
			d1 := g.MinDistance(p, dest, distance)
			m := d + d1
			if m < distance {
				distance = m
			}
		}
	}

	return distance
}
