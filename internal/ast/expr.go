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

func NewStringExpre(value string) *StringExpr {
	return &StringExpr{Value: value}
}

func (s *StringExpr) exprNode() {}

type VariableExpr struct {
	Type string
	Name string
}

func NewVariableExpr(_type, name string) *VariableExpr {
	return &VariableExpr{
		Type: _type,
		Name: name,
	}
}

func (v *VariableExpr) exprNode() {}

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
