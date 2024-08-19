package ast

type ReturnStmt struct {
	Type string
}

func (r *ReturnStmt) stmtNode() {}

type BadStmt struct{}

func (b *BadStmt) stmtNode() {}

type ExprStmt struct {
	X Expr
}

func NewExprStmt(x Expr) *ExprStmt {
	return &ExprStmt{X: x}
}

func (e *ExprStmt) stmtNode() {}
