package ast

import (
	"JATL/token"
	"bytes"
)

type Node interface {
  TokenLiteral() string
  String()       string
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

func (p *Program) String() string {
  var out bytes.Buffer

  for _, s := range p.Statement {
    out.WriteString(s.String())
  }

  return out.String()
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

func (i *Identifier) String() string {
  return i.Value
}

