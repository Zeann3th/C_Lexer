package parser

import (
	"fmt"

	"github.com/zeann3th/C_Compiler/internal/ast"
	lx "github.com/zeann3th/C_Compiler/internal/lexer"
)

type Parser struct {
	*lx.Lexer
	Current  lx.Token
	Previous lx.Token
}

func NewParser(l *lx.Lexer) *Parser {
	return &Parser{
		Lexer: l,
	}
}

func (p *Parser) GetNextToken() {
	p.Previous = p.Current
	p.Current = *p.NextToken()
}

func (p *Parser) ExpectToken(a lx.TokenKind, b ...lx.TokenKind) bool {
	tmp := ""
	for _, kind := range b {
		tmp += lx.Codex[kind]
		tmp += ", "
	}
	msg := fmt.Errorf("Line %v, col %v: ERROR: Expected %vgot %v instead\n", p.Line, p.Col, tmp, lx.Codex[a])
	for _, kind := range b {
		if a == kind {
			return true
		}
	}
	fmt.Println(msg)
	return false
}

func (p *Parser) ParsePrimary() ast.Node {
	switch p.Current.Kind {
	default:
		fmt.Printf("Line %v, col %v: ERROR: Unknown token\n", p.Line, p.Col)
		return ast.BadStmt{}
	case lx.TYPE:
		return p.ParseIdentifierExpr()
	case lx.OCURLY:
		return p.ParseBlock()
	case lx.PREPROCESSOR:
		return ast.BadExpr{}
	}
}
