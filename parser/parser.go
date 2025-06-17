package parser

import (
	"JATL/ast"
	"JATL/lexer"
	"JATL/token"
	"fmt"
)

type (
  prefixParseFn func() ast.Expression
  infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
  l         *lexer.Lexer
  curToken  token.Token
  peekToken token.Token
  errors    []string

  prefixParseFns map[token.TokenType]prefixParseFn
  infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser{
  p := &Parser{l : l, errors: []string{}}
  // Read two tokens so that cur and peek are both set
  p.NextToken()
  p.NextToken()

  p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
  p.infixParseFns = make(map[token.TokenType]infixParseFn)

  p.registerPrefix(token.IDENT , p.parseIdentifier)
  p.registerPrefix(token.INT   , p.parseIntegerLiteral)
  p.registerPrefix(token.BANG  , p.parsePrefixExpression)
  p.registerPrefix(token.MINUS , p.parsePrefixExpression)


  p.registerInfix(token.PLUS       , p.parseInfixExpression)
  p.registerInfix(token.MINUS      , p.parseInfixExpression)
  p.registerInfix(token.ASTERISK   , p.parseInfixExpression)
  p.registerInfix(token.SLASH      , p.parseInfixExpression)
  p.registerInfix(token.GREATER    , p.parseInfixExpression)
  p.registerInfix(token.LESS       , p.parseInfixExpression)
  p.registerInfix(token.EQUAL      , p.parseInfixExpression)
  p.registerInfix(token.BANG_EQUAL , p.parseInfixExpression)

  return p
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn){
  p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn){
  p.infixParseFns[tokenType] = fn
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
    case token.RETURN: {
      return p.parseReturnStatement()
    }
    default:
      return p.parseExpressionStatement()
  }
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

