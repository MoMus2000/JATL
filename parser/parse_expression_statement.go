package parser

import (
	"JATL/ast"
	"strconv"
)

func (p *Parser) parseExpressionStatement() ast.Statement{
  stmt := &ast.ExpresssionStatement{
    Token: p.curToken,
  }
  expr := p.parseExpression(LOWEST)
  stmt.Expression = expr
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

func (p *Parser) parseExpression(_ int) ast.Expression {
  prefix := p.prefixParseFns[p.curToken.Type]

  if prefix == nil {
    return nil
  }

  leftExpression := prefix()

  return leftExpression
}

