package graph

import (
	"math"
	"testing"
)

var graphNodes = []node{
	{0, 0}, {1, 0}, {2, 0}, {3, 0},
	/*   */ {1, 1} /*   */, {3, 1},
	{0, 2}, {1, 2}, /*           */
	/*   */ {1, 3}, {2, 3}, {3, 3},
}

type node struct{ x, y float64 }

func (n node) Expand() []node {
	candidates := []node{
		{n.x + 0, n.y + 1},
		{n.x + 0, n.y - 1},
		{n.x + 1, n.y + 0},
		{n.x - 1, n.y + 0},
	}
	actual := []node{}

	for _, c := range candidates {
		for _, g := range graphNodes {
			if c == g {
				actual = append(actual, c)
			}
		}
	}

	return actual
}

func (a node) Distance(b node) float64 {
	x2 := math.Pow(a.x-b.x, 2)
	y2 := math.Pow(a.x-b.x, 2)
	return math.Sqrt(x2 + y2)
}

func (n node) Huristic(target node) float64 {
	return n.Distance(target)
}

func TestAStart(t *testing.T) {
	start, end := node{2, 0}, node{3, 3}
	expect := []node{{2, 0}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {2, 3}, {3, 3}}

	path, err := AStar(start, end)

	if err != nil {
		t.Fatal("No path found")
	}
	if len(path) != len(expect) {
		t.Fatalf("Path error, expected %v got %v", expect, path)
	}
	for i := range expect {
		if path[i] != expect[i] {
			t.Errorf("Path error at %v, expected %v got %v", i, expect[i], path[i])
		}
	}
}
