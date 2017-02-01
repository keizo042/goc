package ast

type NodeType int

type Node interface {
	Node() []Node

	IsLeaf() bool
	IsNode() bool
	Types() NodeType
}

type Expr interface{}

type Digit int

type String string
