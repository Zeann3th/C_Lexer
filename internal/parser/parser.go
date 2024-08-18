package parser

import (
	"fmt"

	ast "github.com/zeann3th/C_Compiler/internal/ast"
	lx "github.com/zeann3th/C_Compiler/internal/lexer"
)

type Parser struct {
	*lx.Lexer
	Current lx.Token
}

func NewParser(l lx.Lexer) *Parser {
	return &Parser{
		Lexer: &l,
	}
}

func (p *Parser) GetNextToken() {
	p.Current = *p.NextToken()
}

func (p *Parser) ParseExpr() ast.Expr

func (p *Parser) ParseNumberExpr() ast.Expr {
	result := ast.NewNumberExpr(p.NumBuf)
	p.NextToken()
	return result
}

func (p *Parser) ParseParenExpr() ast.Expr {
	p.NextToken()
	v := p.ParseExpr()
	if p.Current.Kind != lx.CPAREN {
		fmt.Printf("Line %v, col %v: ERROR: Expected %v but got %v instead\n", p.Line, p.Col, lx.Codex[lx.CPAREN], lx.Codex[p.Current.Kind])
	}
	return v
}

func (p *Parser) ParsePrimary() ast.Expr {
	switch p.Current.Kind {
	default:
		fmt.Printf("Line %v, col %v: ERROR: Unknown token", p.Line, p.Col)
		return nil
	case lx.NUMBER:
		return p.ParseNumberExpr()
	case lx.OPAREN:
		return p.ParseParenExpr()
	}
}
