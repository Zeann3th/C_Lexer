package parser

import "github.com/zeann3th/C_Compiler/internal/ast"

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

func (p *Parser) ParseAssignExpr() ast.Stmt {
	return &ast.BadStmt{}
}

func (p *Parser) ParseFnCallExpr() ast.Stmt {
	return &ast.BadStmt{}
}

func (p *Parser) ParseIdentifierExpr() ast.Stmt {
	return &ast.BadStmt{}
}
