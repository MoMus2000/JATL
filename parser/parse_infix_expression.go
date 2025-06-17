package parser

import "JATL/ast"

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
  stmt := &ast.InfixExpression{
    Token: p.curToken,
    Left: left,
  }

  precendence := p.curPrecedence()

  p.NextToken()

  expression  := p.parseExpression(precendence)

  stmt.Right = expression

  return stmt
}

