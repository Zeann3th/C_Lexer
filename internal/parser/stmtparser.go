package parser

import (
	ast "github.com/zeann3th/C_Compiler/internal/ast"
	lx "github.com/zeann3th/C_Compiler/internal/lexer"
)

func (p *Parser) ParseBlock() []ast.Stmt {
	block := []ast.Stmt{}
	var stmt ast.Stmt
	p.GetNextToken() // Eat '{'
	for {
		switch p.Current.Kind {
		case lx.TYPE:
			stmt = p.ParseIdentifierExpr() // Check Declarations, assigments
			block = append(block, stmt)
		case lx.KEYWORD:
			stmt = p.ParseKeywordStmt()
			block = append(block, stmt)
		case lx.SYMBOL:
			p.GetNextToken()
			switch p.Current.Kind {
			case lx.OPAREN:
				stmt = p.ParseFnCallExpr()
				block = append(block, stmt)
			case lx.EQ:
				stmt = p.ParseAssignExpr()
				block = append(block, stmt)
			}
		case lx.RETURN:
			p.GetNextToken()
			// ExpectToken from the previously appended type
		case lx.CCURLY:
			return block
		default:
			p.ExpectToken(p.Current.Kind, lx.TYPE, lx.KEYWORD, lx.SYMBOL, lx.CCURLY)
			block = append(block, &ast.BadStmt{})
		}
	}
}

func (p *Parser) ParseKeywordStmt() ast.Stmt {
	return &ast.BadStmt{}
}
