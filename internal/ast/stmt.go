package ast

type BlockStmt struct {
	Stmts []Stmt
}

func (b *BlockStmt) stmtNode() {}

type ReturnStmt struct {
	Value Expr
}

func (r *ReturnStmt) stmtNode() {}

type IfStmt struct {
	Condition Expr
	Then      Stmt
	Else      Stmt
}

func (i *IfStmt) stmtNode() {}
