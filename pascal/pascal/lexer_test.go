package pascal

import "testing"

func TestLexerIntegers(t *testing.T) {
	lexer := NewLexer("123 456 789")
	
	token, err := lexer.NextToken()
	if err != nil || token.Type != INTEGER || token.Value != "123" {
		t.Errorf("expected INTEGER 123, got %v, error: %v", token, err)
	}
	
	token, err = lexer.NextToken()
	if err != nil || token.Type != INTEGER || token.Value != "456" {
		t.Errorf("expected INTEGER 456, got %v, error: %v", token, err)
	}
	
	token, err = lexer.NextToken()
	if err != nil || token.Type != INTEGER || token.Value != "789" {
		t.Errorf("expected INTEGER 789, got %v, error: %v", token, err)
	}
	
	token, err = lexer.NextToken()
	if err != nil || token.Type != EOF {
		t.Errorf("expected EOF, got %v, error: %v", token, err)
	}
}

func TestLexerOperators(t *testing.T) {
	lexer := NewLexer("+ - * /")
	
	expected := []TokenType{PLUS, MINUS, MUL, DIV, EOF}
	
	for _, expectedType := range expected {
		token, err := lexer.NextToken()
		if err != nil || token.Type != expectedType {
			t.Errorf("expected %d, got %v, error: %v", expectedType, token, err)
		}
	}
}

func TestLexerParentheses(t *testing.T) {
	lexer := NewLexer("(2+3)")
	
	expected := []TokenType{LPAREN, INTEGER, PLUS, INTEGER, RPAREN, EOF}
	
	for _, expectedType := range expected {
		token, err := lexer.NextToken()
		if err != nil || token.Type != expectedType {
			t.Errorf("expected %d, got %v, error: %v", expectedType, token, err)
		}
	}
}

func TestLexerKeywords(t *testing.T) {
	lexer := NewLexer("BEGIN END begin end Begin End")
	
	expected := []TokenType{BEGIN, END, BEGIN, END, BEGIN, END, EOF}
	
	for _, expectedType := range expected {
		token, err := lexer.NextToken()
		if err != nil || token.Type != expectedType {
			t.Errorf("expected %d, got %v, error: %v", expectedType, token, err)
		}
	}
}

func TestLexerIdentifiers(t *testing.T) {
	lexer := NewLexer("x y abc test123 _var")
	
	expected := []string{"x", "y", "abc", "test123", "_var"}
	
	for _, expectedValue := range expected {
		token, err := lexer.NextToken()
		if err != nil || token.Type != ID || token.Value != expectedValue {
			t.Errorf("expected ID %s, got %v, error: %v", expectedValue, token, err)
		}
	}
}

func TestLexerAssignment(t *testing.T) {
	lexer := NewLexer("x := 5")
	
	token, err := lexer.NextToken()
	if err != nil || token.Type != ID || token.Value != "x" {
		t.Errorf("expected ID x, got %v, error: %v", token, err)
	}
	
	token, err = lexer.NextToken()
	if err != nil || token.Type != ASSIGN {
		t.Errorf("expected ASSIGN, got %v, error: %v", token, err)
	}
	
	token, err = lexer.NextToken()
	if err != nil || token.Type != INTEGER || token.Value != "5" {
		t.Errorf("expected INTEGER 5, got %v, error: %v", token, err)
	}
}

func TestLexerSemicolon(t *testing.T) {
	lexer := NewLexer("x := 5; y := 10")
	
	expected := []TokenType{ID, ASSIGN, INTEGER, SEMI, ID, ASSIGN, INTEGER, EOF}
	
	for _, expectedType := range expected {
		token, err := lexer.NextToken()
		if err != nil || token.Type != expectedType {
			t.Errorf("expected %d, got %v, error: %v", expectedType, token, err)
		}
	}
}

func TestLexerDot(t *testing.T) {
	lexer := NewLexer("END.")
	
	token, err := lexer.NextToken()
	if err != nil || token.Type != END {
		t.Errorf("expected END, got %v, error: %v", token, err)
	}
	
	token, err = lexer.NextToken()
	if err != nil || token.Type != DOT {
		t.Errorf("expected DOT (.), got %v, error: %v", token, err)
	}
}

func TestLexerWhitespace(t *testing.T) {
	lexer := NewLexer("  1   +   2  ")
	
	token, err := lexer.NextToken()
	if err != nil || token.Type != INTEGER || token.Value != "1" {
		t.Errorf("expected INTEGER 1, got %v, error: %v", token, err)
	}
	
	token, err = lexer.NextToken()
	if err != nil || token.Type != PLUS {
		t.Errorf("expected PLUS, got %v, error: %v", token, err)
	}
	
	token, err = lexer.NextToken()
	if err != nil || token.Type != INTEGER || token.Value != "2" {
		t.Errorf("expected INTEGER 2, got %v, error: %v", token, err)
	}
}

func TestLexerInvalidCharacter(t *testing.T) {
	lexer := NewLexer("@")
	
	_, err := lexer.NextToken()
	if err == nil {
		t.Error("expected error for invalid character")
	}
}