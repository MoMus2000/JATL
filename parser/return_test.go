package parser

import (
	"JATL/ast"
	"JATL/lexer"
	"testing"
)


func TestReturnStatement(t *testing.T){
  input := `
    return 5;
    return 10;
    return 993322;
  `

  l := lexer.New(input)
  p := New(l)
  program := p.ParseProgram()
  checkParserError(t, p)


  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }

  if len(program.Statement) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. Got %d",
      len(program.Statement),
    )
  }

  for _, statement := range program.Statement{

    returnstmt, ok := statement.(*ast.ReturnStatement)

    if !ok {
      t.Errorf("stmt not *ast.ReturnStatement, got %T", statement)
      continue
    }

    if returnstmt.TokenLiteral() != "return" {
      t.Errorf("returnStmt.Token.Literal not `return`, got %s", returnstmt.TokenLiteral())
    }
  }

}
