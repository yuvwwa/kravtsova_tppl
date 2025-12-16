package pascal

import "testing"

func TestParserEmptyProgram(t *testing.T) {
	text := "BEGIN END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if _, ok := node.(Compound); !ok {
		t.Errorf("ожидался узел Compound, получен %T", node)
	}
}

func TestParserSimpleAssignment(t *testing.T) {
	text := "BEGIN x := 5 END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	compound, ok := node.(Compound)
	if !ok {
		t.Fatalf("ожидался узел Compound, получен %T", node)
	}
	
	if len(compound.Children) != 1 {
		t.Errorf("ожидался 1 дочерний узел, получено %d", len(compound.Children))
	}
	
	if _, ok := compound.Children[0].(Assign); !ok {
		t.Errorf("ожидался узел Assign, получен %T", compound.Children[0])
	}
}

func TestParserMultipleAssignments(t *testing.T) {
	text := "BEGIN x := 5; y := 10 END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	compound, ok := node.(Compound)
	if !ok {
		t.Fatalf("ожидался узел Compound, получен %T", node)
	}
	
	if len(compound.Children) != 2 {
		t.Errorf("ожидалось 2 дочерних узла, получено %d", len(compound.Children))
	}
}

func TestParserNestedCompound(t *testing.T) {
	text := "BEGIN BEGIN x := 5 END END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	compound, ok := node.(Compound)
	if !ok {
		t.Fatalf("ожидался узел Compound, получен %T", node)
	}
	
	if len(compound.Children) != 1 {
		t.Errorf("ожидался 1 дочерний узел, получено %d", len(compound.Children))
	}
	
	if _, ok := compound.Children[0].(Compound); !ok {
		t.Errorf("ожидался вложенный узел Compound, получен %T", compound.Children[0])
	}
}

func TestParserArithmeticExpression(t *testing.T) {
	text := "BEGIN x := 2 + 3 * 4 END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	_, err = parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParserParentheses(t *testing.T) {
	text := "BEGIN x := (2 + 3) * 4 END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	_, err = parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParserUnaryOperators(t *testing.T) {
	text := "BEGIN x := -5; y := +10 END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	_, err = parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParserVariableInExpression(t *testing.T) {
	text := "BEGIN x := 5; y := x + 1 END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	_, err = parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParserEmptyStatements(t *testing.T) {
	text := "BEGIN ; ; x := 5; ; END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	_, err = parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParserMissingDot(t *testing.T) {
	text := "BEGIN x := 5 END"
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	_, err = parser.Parse()
	if err == nil {
		t.Error("ожидалась ошибка из-за отсутствия точки в конце программы")
	}
}

func TestParserInvalidSyntax(t *testing.T) {
	text := "BEGIN x := END."
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	_, err = parser.Parse()
	if err == nil {
		t.Error("expected syntax error")
	}
}