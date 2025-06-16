package ast

import "JATL/token"

type Node interface {
  TokenLiteral() string
}

type Statement interface {
  Node
  statementNode()
}

type Expression interface {
  Node
  expressionNode()
}

type Program struct {
  Statement []Statement
}

func (p *Program) TokenLiteral() string {
  if len(p.Statement) > 0 {
    return p.Statement[0].TokenLiteral()
  }
  return ""
}

type Identifier struct {
  Token token.Token
  Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
  return i.Token.Literal
}

