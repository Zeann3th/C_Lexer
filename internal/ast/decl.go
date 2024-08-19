package ast

type VarDecl struct {
	Type  string
	Name  string
	Value Expr
}

func NewVarDecl(_type, name string, value Expr) *VarDecl {
	return &VarDecl{
		Type:  _type,
		Name:  name,
		Value: value,
	}
}

func (v *VarDecl) declNode() {}
func (v *VarDecl) stmtNode() {}

type FuncDecl struct {
	Name   string
	Params []*VarDecl
	Body   []Stmt
}

func NewFuncDecl(name string, params []*VarDecl, body []Stmt) *FuncDecl {
	return &FuncDecl{
		Name:   name,
		Params: params,
		Body:   body,
	}
}

func (fd *FuncDecl) declNode() {}
func (fd *FuncDecl) stmtNode() {}
