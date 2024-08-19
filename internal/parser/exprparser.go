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
		return nil
	}
	p.GetNextToken()
	if p.ExpectToken(p.Current.Kind, lx.SEMICOLON) {
		return ast.NewAssignExpr(left, right)
	}
	return nil
}

func (p *Parser) ParseFnCallExpr() ast.Expr {
	return nil
}

func (p *Parser) ParseIdentifierExpr() ast.Expr {
	returnType, ok := p.Current.Value.(string)
	if !ok {
		return nil
	}
	typeHandlers := map[string]lx.TokenKind{
		"string": lx.STRING,
		"int":    lx.NUMBER,
		"float":  lx.NUMBER,
		"char":   lx.CHAR,
	}
	initValues := map[string]ast.Expr{
		"string": ast.NewStringExpr(""),
		"int":    ast.NewNumberExpr(0),
		"float":  ast.NewNumberExpr(0.0),
		"char":   ast.NewStringExpr(""),
	}
	name := ""
	p.GetNextToken()
	if p.ExpectToken(p.Current.Kind, lx.SYMBOL) {
		name = p.Current.Name
		p.GetNextToken()
		switch p.Current.Kind {
		case lx.OPAREN:
			return p.ParseFnCallExpr()
		case lx.ASSIGN:
			p.GetNextToken()
			if p.ExpectToken(p.Current.Kind, typeHandlers[returnType]) {
				leftInit := ast.NewVarDecl(returnType, name, initValues[returnType])
				left := ast.NewVarDeclExpr(leftInit)
				p.GetNextToken()
				if p.ExpectToken(p.Current.Kind, lx.SEMICOLON) {
					return p.ParseAssignExpr(left)
				}
				return nil
			}
		case lx.SEMICOLON:
			leftInit := ast.NewVarDecl(returnType, name, initValues[returnType])
			return ast.NewVarDeclExpr(leftInit)
		default:
			p.ExpectToken(p.Current.Kind, lx.OPAREN, lx.ASSIGN)
			return nil
		}
	}
	return nil
}
