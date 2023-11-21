package graph

type Node[TNode any] interface {
	comparable
	Expandable[TNode]
	Measurable[TNode]
}

type Expandable[TNode any] interface {
	Expand() []TNode
}

type Measurable[TNode any] interface {
	Distance(to TNode) float64
}
