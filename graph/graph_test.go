package graph

import "testing"

func Test_Graph(t *testing.T) {
	graph := Graph{
		Points: make(map[*Point]struct{}),
	}
	p1 := Point{
		Name:  "1",
		Edges: []*Edge{},
	}

	p2 := Point{
		Name:  "2",
		Edges: []*Edge{},
	}

	p3 := Point{
		Name:  "3",
		Edges: []*Edge{},
	}

	p4 := Point{
		Name:  "4",
		Edges: []*Edge{},
	}

	p5 := Point{
		Name:  "5",
		Edges: []*Edge{},
	}

	p6 := Point{
		Name:  "6",
		Edges: []*Edge{},
	}

	p7 := Point{
		Name:  "7",
		Edges: []*Edge{},
	}

	p1.AddEdge(&p3, 10)
	p1.AddEdge(&p5, 30)
	p1.AddEdge(&p6, 100)
	p2.AddEdge(&p3, 5)
	p3.AddEdge(&p4, 50)
	p4.AddEdge(&p6, 10)
	p5.AddEdge(&p4, 20)
	p5.AddEdge(&p6, 60)
	p6.AddEdge(&p7, 90000)

	graph.Points[&p1] = struct{}{}
	graph.Points[&p2] = struct{}{}
	graph.Points[&p3] = struct{}{}
	graph.Points[&p4] = struct{}{}
	graph.Points[&p5] = struct{}{}
	graph.Points[&p6] = struct{}{}
	graph.Points[&p7] = struct{}{}

	if !graph.HasPoint(&p1) {
		t.Fail()
	}

	graph.BuildMinimalDistance()
	t.Errorf("%v\n", graph.MinimalDistances)

	md := graph.MinDistance(&p1, &p2, MaxInt)
	md1 := graph.MinDistance(&p1, &p4, MaxInt)
	md2 := graph.MinDistance(&p1, &p5, MaxInt)
	md3 := graph.MinDistance(&p2, &p6, MaxInt)
	md4 := graph.MinDistance(&p2, &p7, MaxInt)
	t.Errorf("P1 --> P2 MinDistance: %v\n", md)
	t.Errorf("P1 --> P4 MinDistance: %v\n", md1)
	t.Errorf("P1 --> P5 MinDistance: %v\n", md2)
	t.Errorf("P2 --> P6 MinDistance: %v\n", md3)
	t.Errorf("P2 --> P7 MinDistance: %v\n", md4)
}
