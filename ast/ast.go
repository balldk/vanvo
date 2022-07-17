package ast

type Node interface {
	String() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}
