package ast

import "JATL/token"

type ReturnStatement struct {
  Token       token.Token
  ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
  return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {

}
