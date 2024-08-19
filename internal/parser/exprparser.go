package parser

import (
	"github.com/zeann3th/C_Compiler/internal/ast"
	lx "github.com/zeann3th/C_Compiler/internal/lexer"
)

func (p *Parser) ParseNumberExpr() ast.Expr {
	result := ast.NewNumberExpr(p.NumBuf)
	p.NextToken()
	return result
}

func (p *Parser) ParseStringExpr() ast.Expr {
	result := ast.NewStringExpr(p.Bufnr)
	p.NextToken()
	return result
}

func (p *Parser) ParseAssignExpr(left ast.Expr) ast.Expr {
	p.GetNextToken()
	var right ast.Expr
	switch p.Current.Kind {
	case lx.NUMBER:
		right = p.ParseNumberExpr()
	case lx.STRING:
		right = p.ParseStringExpr()
	case lx.SYMBOL:
	default:
		p.ExpectToken(p.Current.Kind, lx.NUMBER, lx.STRING, lx.SYMBOL)
		return &ast.BadExpr{}
	}
	p.GetNextToken()
	if p.ExpectToken(p.Current.Kind, lx.SEMICOLON) {
		return ast.NewAssignExpr(left, right)
	}
	return &ast.BadExpr{}
}

func (p *Parser) ParseFnCallExpr() ast.Expr {
	return &ast.BadExpr{}
}

func (p *Parser) ParseIdentifierExpr() ast.Expr {
	return &ast.BadExpr{}
}
