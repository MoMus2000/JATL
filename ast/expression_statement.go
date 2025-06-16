package ast

import (
	"JATL/token"
)

type ExpresssionStatement struct {
  Token      token.Token
  Expression Expression
}

func (es *ExpresssionStatement) statementNode() {

}

func (es *ExpresssionStatement) TokenLiteral() string{
  return es.Token.Literal
}

func (es *ExpresssionStatement) String() string {
  if es.Expression != nil {
    return es.Expression.String()
  }
  return ""
}
