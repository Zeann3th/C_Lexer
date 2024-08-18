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
	p.GetNextToken()
	v := p.ParseExpr()
	if v == nil {
		return nil
	}
	if p.Current.Kind != lx.CPAREN {
		fmt.Printf("Line %v, col %v: ERROR: Expected %v but got %v instead\n", p.Line, p.Col, lx.Codex[lx.CPAREN], lx.Codex[p.Current.Kind])
	}
	return v
}

func (p *Parser) ParseIdentifierExpr() ast.Expr { // Needs more checking
	// Return type
	_type := p.Bufnr
	p.GetNextToken()
	// Identifier
	name := ""
	if p.Current.Kind == lx.SYMBOL {
		name = p.Bufnr
	} else {
		fmt.Printf("Line %v, col %v: ERROR: Expected %v but got %v instead\n", p.Line, p.Col, lx.Codex[lx.SYMBOL], lx.Codex[p.Current.Kind])
		return nil
	}
	p.GetNextToken()
	if p.Current.Kind != lx.OPAREN {
		return ast.NewVariableExpr(_type, name)
	}
	p.GetNextToken()
	args := []ast.Expr{}
	if p.Current.Kind != lx.CPAREN {
		for {
			if arg := p.ParseExpr(); arg != nil {
				args = append(args, arg)
			} else {
				return nil
			}
			if p.Current.Kind == lx.CPAREN {
				break
			}
			if p.Current.Kind != lx.COMMA {
				fmt.Printf("Line %v, col %v: ERROR: Expected %v or %v but got %v instead\n", p.Line, p.Col, lx.Codex[lx.CPAREN], lx.Codex[lx.COMMA], lx.Codex[p.Current.Kind])
				return nil
			}
			p.GetNextToken()
		}
	}
	p.GetNextToken()
	return ast.NewCallExpr(_type, name, args)
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
	case lx.TYPE:
		return p.ParseIdentifierExpr()
	}
}
