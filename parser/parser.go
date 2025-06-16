package parser

import (
	"JATL/ast"
	"JATL/lexer"
	"JATL/token"
	"fmt"
)

type Parser struct {
  l         *lexer.Lexer
  curToken  token.Token
  peekToken token.Token
  errors    []string
}

func New(l *lexer.Lexer) *Parser{
  p := &Parser{l : l, errors: []string{}}
  // Read two tokens so that cur and peek are both set
  p.NextToken()
  p.NextToken()
  return p
}

func (p *Parser) Errors() []string{
  return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
  msg := fmt.Sprintf("expected next token to be %s but got %s", t, p.curToken.Type)
  p.errors = append(p.errors, msg)
}

func (p *Parser) NextToken() {
  p.curToken  = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
  program := &ast.Program{}
  program.Statement = []ast.Statement{}
  for p.curToken.Type != token.EOF {
    stmt := p.parseStatement()
    if stmt != nil {
      program.Statement = append(program.Statement, stmt)
    }
    p.NextToken()
  }
  return program
}

func (p *Parser) parseStatement() ast.Statement {
  switch p.curToken.Type {
    case token.LET: {
      return p.parseLetStatement()
    }
    default:
      return nil
  }
}

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

func (p *Parser) expectPeek(tokenType token.TokenType) bool{
  if p.peekTokenIs(tokenType){
    p.NextToken()
    return true
  }
  p.peekError(tokenType)
  return false
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
  return t == p.curToken.Type
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
  return p.peekToken.Type == t
}

