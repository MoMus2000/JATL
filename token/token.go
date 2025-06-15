package token

type TokenType string

const (
  ILLEGAL    = "ILLEGAL"
  EOF        = "EOF"

  IDENT      = "IDENT"
  INT        = "INT"

  ASSIGN     = "="
  PLUS       = "+"
  MINUS      = "-"
  BANG       = "!"
  ASTERISK   = "*"
  LESS       = "<"
  GREATER    = ">"
  SLASH      = "/"

  COMMA      = ","
  SEMICOLON  = ";"

  LPAREN     = "("
  RPAREN     = ")"

  LBRACE     = "{"
  RBRACE     = "}"

  FUNCTION   = "FUNCTION"
  LET        = "LET"
)

var keywords = map[string]TokenType {
  "fn" : FUNCTION,
  "let": LET,
}

func LookupIdent(identString string) TokenType{
  keyword, found := keywords[identString]

  if !found {
    return IDENT
  }

  return keyword
}

type Token struct {
  Type       TokenType
  Literal    string
}

