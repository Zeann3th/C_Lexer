package ast

type VarDecl struct {
	Name  string
	Value Expr
}

func (v *VarDecl) declNode() {}

type FuncDecl struct {
	Name   string
	Params []*VarDecl
	Body   []Stmt
}

func (fd *FuncDecl) declNode() {}
