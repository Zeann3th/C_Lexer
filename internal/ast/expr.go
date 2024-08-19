package ast

type NumberExpr struct {
	Value float64
}

func NewNumberExpr(value float64) *NumberExpr {
	return &NumberExpr{Value: value}
}

func (n *NumberExpr) exprNode() {}

type StringExpr struct {
	Value string
}

func NewStringExpr(value string) *StringExpr {
	return &StringExpr{Value: value}
}

func (s *StringExpr) exprNode() {}

type VarExpr struct {
	Type string
	Name string
}

func NewVarExpr(_type, name string) *VarExpr {
	return &VarExpr{
		Type: _type,
		Name: name,
	}
}

func NewVarDeclExpr(decl *VarDecl) *VarExpr {
	return &VarExpr{
		Type: decl.Type,
		Name: decl.Name,
	}
}

func (v *VarExpr) exprNode() {}

type BinaryExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (b *BinaryExpr) exprNode() {}

type CallExpr struct {
	Type   string
	Callee string
	Args   []Expr
}

func NewCallExpr(_type, callee string, args []Expr) *CallExpr {
	return &CallExpr{
		Type:   _type,
		Callee: callee,
		Args:   args,
	}
}

func (c *CallExpr) exprNode() {}

type AssignExpr struct {
	Left  Expr
	Right Expr
}

func NewAssignExpr(left, right Expr) *AssignExpr {
	return &AssignExpr{
		Left:  left,
		Right: right,
	}
}

func (a *AssignExpr) exprNode() {}

type BadExpr struct{}

func (b *BadExpr) exprNode() {}
