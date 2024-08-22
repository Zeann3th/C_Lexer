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
	case lx.VAR:
	default:
		p.ExpectToken(p.Current.Kind, lx.NUMBER, lx.STRING, lx.VAR)
		return nil
	}
	p.GetNextToken()
	if p.ExpectToken(p.Current.Kind, lx.SEMICOLON) {
		return ast.NewAssignExpr(left, right)
	}
	return nil
}

func (p *Parser) ParseFnCallExpr() ast.Expr {
	p.GetNextToken()
	for {
		switch p.Current.Kind {
		case lx.FUNC:
		case lx.CPAREN:
		case lx.STRING:
		default:
			p.ExpectToken(p.Current.Kind, lx.FUNC, lx.CPAREN, lx.STRING)
			return nil
		}
	}
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
	p.GetNextToken()
	name, ok := p.Current.Value.(string)
	if !ok {
		return nil
	}
	switch p.Current.Kind {
	case lx.VAR:
		varValue := initValues[returnType]
		for {
			p.GetNextToken()
			switch p.Current.Kind {
			case lx.ASSIGN:
				p.GetNextToken()
				if !p.ExpectToken(p.Current.Kind, typeHandlers[returnType]) {
					return nil
				}
				switch p.Current.Kind {
				case lx.STRING:
					tmp, ok := p.Current.Value.(string)
					if !ok {
						return nil
					}
					varValue = ast.NewStringExpr(tmp)
				case lx.NUMBER:
					tmp, ok := p.Current.Value.(float64)
					if !ok {
						return nil
					}
					varValue = ast.NewNumberExpr(tmp)
				}
			case lx.SEMICOLON:
				return ast.NewDeclExpr(ast.NewVarDecl(returnType, name, varValue))
			default:
				p.ExpectToken(p.Current.Kind, lx.ASSIGN, lx.SEMICOLON)
				return nil
			}
		}
	case lx.FUNC:
		p.GetNextToken()
		if !p.ExpectToken(p.Current.Kind, lx.OPAREN) {
			return nil
		}
		p.GetNextToken()
		return nil
	default:
		p.ExpectToken(p.Current.Kind, lx.VAR, lx.FUNC)
		return nil
	}
}
