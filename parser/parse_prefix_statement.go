package parser

import "JATL/ast"


func (p *Parser) parsePrefixExpression() ast.Expression{

  expr := &ast.PrefixExpression{
    Token   : p.curToken,
    Operator: p.curToken.Literal,
  }

  p.NextToken()

  expr.Right = p.parseExpression(PREFIX)
  
  return expr
}

