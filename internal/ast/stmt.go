package ast

type ReturnStmt struct{}

func (r *ReturnStmt) stmtNode() {}

type BadStmt struct{}

func (b *BadStmt) stmtNode() {}
