package ast

type NodeType int

type Node interface {
	Node() []Node

	IsLeaf() bool
	IsNode() bool
	Types() NodeType
}

type Expr interface {
	Node
	exprNode()
}

type exprNode struct {
}

func (e *exprNode) exprNode() {}

type BinOpExpr struct {
	exprNode
	Lhs Expr
	Op  Item
	Rhs Expr
}
