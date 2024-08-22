package parser

import (
	ast "github.com/zeann3th/C_Compiler/internal/ast"
	lx "github.com/zeann3th/C_Compiler/internal/lexer"
)

func (p *Parser) ParseBlock() []ast.Stmt {
	var block []ast.Stmt
	var stmt ast.Stmt
	p.GetNextToken() // Eat '{'
	for {
		switch p.Current.Kind {
		case lx.TYPE:
			stmt = ast.NewExprStmt(p.ParseIdentifierExpr())
			block = append(block, stmt)
		case lx.KEYWORD:
			stmt = p.ParseKeywordStmt()
			block = append(block, stmt)
		case lx.FUNC:
			p.GetNextToken()
			switch p.Current.Kind {
			case lx.OPAREN:
				stmt = ast.NewExprStmt(p.ParseFnCallExpr())
				block = append(block, stmt)
			case lx.ASSIGN:
				left := p.ParseNumberExpr()
				stmt = ast.NewExprStmt(p.ParseAssignExpr(left))
				block = append(block, stmt)
			}
		case lx.RETURN:
			p.GetNextToken()
			stmt = p.ParseReturnStmt()
			block = append(block, stmt)
			// ExpectToken from the previously appended type
		case lx.CCURLY:
			return block
		default:
			p.ExpectToken(p.Current.Kind, lx.TYPE, lx.KEYWORD, lx.FUNC, lx.CCURLY)
			block = append(block, nil)
		}
		p.GetNextToken()
	}
}

func (p *Parser) ParseReturnStmt() ast.Stmt {
	return nil
}

func (p *Parser) ParseKeywordStmt() ast.Stmt {
	return nil
}
