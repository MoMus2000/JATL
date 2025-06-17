package parser

import (
	"JATL/ast"
	"JATL/token"
	"strconv"
)

func (p *Parser) parseExpressionStatement() ast.Statement{
  stmt := &ast.ExpresssionStatement{
    Token: p.curToken,
  }
  expr := p.parseExpression(LOWEST)

  stmt.Expression = expr

  if p.peekTokenIs(token.SEMICOLON){
    p.NextToken()
  }

  return  stmt
}

func (p *Parser) parseIdentifier() ast.Expression{
    return &ast.Identifier{
      Token: p.curToken,
      Value: p.curToken.Literal,
  }
}

func (p *Parser) parseIntegerLiteral() ast.Expression{
    intVal, _ := strconv.ParseInt(p.curToken.Literal, 10, 64)
    return &ast.IntegerLiteral{
      Token: p.curToken,
      Value: intVal,
  }
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
  prefix := p.prefixParseFns[p.curToken.Type]

  if prefix == nil {
    return nil
  }

  leftExpression := prefix()

  for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
    infix := p.infixParseFns[p.peekToken.Type]
    if infix == nil {
      return leftExpression
    }

    p.NextToken()
    leftExpression = infix(leftExpression)
  }

  return leftExpression
}

