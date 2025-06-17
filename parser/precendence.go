package parser

import "JATL/token"

const (
  _ int = iota
  LOWEST
  EQUALS
  LESSGREATER
  SUM
  PRODUCT
  PREFIX
  CALL
)

var precendences = map[token.TokenType]int{
  token.EQUAL         : EQUALS,
  token.BANG_EQUAL    : EQUALS,
  token.LESS          : LESSGREATER,
  token.GREATER       : LESSGREATER,
  token.PLUS          : SUM,
  token.MINUS         : SUM,
  token.ASTERISK      : PRODUCT,
  token.SLASH         : PRODUCT,
}

func (p *Parser) peekPrecedence() int {
  pre, ok := precendences[p.peekToken.Type]
  if ok {
    return pre
  }
  return LOWEST
}

func (p *Parser) curPrecedence() int {
  pre, ok := precendences[p.curToken.Type]
  if ok {
    return pre
  }
  return LOWEST
}

