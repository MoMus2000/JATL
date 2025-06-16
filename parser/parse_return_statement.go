package parser

import (
	"JATL/ast"
	"JATL/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement{
  stmt := &ast.ReturnStatement{Token: p.curToken}

  p.NextToken()

  for !p.curTokenIs(token.SEMICOLON){
    p.NextToken()
  }

  return stmt
}

