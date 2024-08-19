package ast

type Node interface{}

type Decl interface {
	Node
	declNode()
}

type Expr interface {
	Node
	exprNode()
}

type Stmt interface {
	Node
	stmtNode()
}

type Program struct {
	Pid  int
	Body []Stmt
}
