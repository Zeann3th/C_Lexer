package parser

type Node struct {
	Type  string
	Left  *Node
	Right *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{
		Left:  left,
		Right: right,
	}
}

type Program struct {
	Body []Stmt
}

type Stmt struct {
	Node
}

type ExprStmt struct {
	Stmt
	Content interface{}
}

type BinaryExpr struct {
	ExprStmt
}
