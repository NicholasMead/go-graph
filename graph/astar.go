package graph

import (
	"errors"
	"fmt"
)

type AStarNode[TNode any] interface {
	Node[TNode]
	Huristic(target TNode) float64
}

type position[TNode AStarNode[TNode]] struct {
	node     TNode
	path     []TNode
	cost     float64
	huristic float64
}

func (p position[TNode]) score() float64 {
	return p.cost + p.huristic
}

func AStar[TNode AStarNode[TNode]](from, to TNode) ([]TNode, error) {
	cache := map[TNode]float64{
		from: 0,
	}
	queue := []position[TNode]{{
		from,
		[]TNode{from},
		0,
		from.Huristic(to),
	}}
	var current position[TNode]

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		fmt.Println(current.node, current.path)

		if current.node == to {
			return current.path, nil
		}

		for _, next := range current.node.Expand() {
			position := position[TNode]{
				next,
				combine(current.path, next),
				current.cost + current.node.Distance(next),
				next.Huristic(to),
			}

			if c, ok := cache[next]; ok && c < position.cost {
				continue
			} else {
				cache[next] = position.cost
			}

			queue = insertInPlace(queue, position)
		}
	}

	return nil, errors.New("no path found")
}

func combine[TNode any](path []TNode, next TNode) []TNode {
	new := make([]TNode, len(path)+1)
	copy(new, path)
	new[len(path)] = next
	return new
}

func insertInPlace[TNode AStarNode[TNode]](queue []position[TNode], insert position[TNode]) []position[TNode] {
	for i, p := range queue {
		if p.score() > insert.score() {
			target := make([]position[TNode], len(queue)+1)
			copy(target, queue[:i])
			copy(target[i+1:], queue[i:])
			target[i] = insert
			return target
		}
	}
	return append(queue, insert)
}
