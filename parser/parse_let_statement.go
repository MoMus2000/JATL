package parser

import (
	"JATL/ast"
	"JATL/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
  stmt := &ast.LetStatement{Token: p.curToken}


  if !p.expectPeek(token.IDENT){
    return nil
  }

  stmt.Name = &ast.Identifier{
    Token: p.curToken,
    Value: p.curToken.Literal,
  }

  if !p.expectPeek(token.ASSIGN){
    return nil
  }

  for !p.curTokenIs(token.SEMICOLON){
    p.NextToken()
  }
  
  return stmt
}
