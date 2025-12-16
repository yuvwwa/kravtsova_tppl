package pascal

import (
	"testing"
)

func TestEmptyProgram(t *testing.T) {
	text := `BEGIN
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if len(vars) != 0 {
		t.Errorf("expected empty variables, got %v", vars)
	}
}

func TestSimpleAssignments(t *testing.T) {
	text := `BEGIN
	x:= 2 + 3 * (2 + 3);
	y:= 2 / 2 - 2 + 3 * ((1 + 1) + (1 + 1))
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	expectedX := 2.0 + 3.0*(2.0+3.0) // 17
	expectedY := 2.0/2.0 - 2.0 + 3.0*((1.0+1.0)+(1.0+1.0)) // 11
	
	if vars["x"] != expectedX {
		t.Errorf("expected x=%f, got %f", expectedX, vars["x"])
	}
	
	if vars["y"] != expectedY {
		t.Errorf("expected y=%f, got %f", expectedY, vars["y"])
	}
}

func TestNestedBlocks(t *testing.T) {
	text := `BEGIN
	y := 2;
	BEGIN
		a := 3;
		a := a;
		b := 10 + a + 10 * y / 4;
		c := a - b
	END;
	x := 11
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	expectedY := 2.0
	expectedA := 3.0
	expectedB := 10.0 + 3.0 + 10.0*2.0/4.0 // 18
	expectedC := 3.0 - 18.0                 // -15
	expectedX := 11.0
	
	if vars["y"] != expectedY {
		t.Errorf("expected y=%f, got %f", expectedY, vars["y"])
	}
	
	if vars["a"] != expectedA {
		t.Errorf("expected a=%f, got %f", expectedA, vars["a"])
	}
	
	if vars["b"] != expectedB {
		t.Errorf("expected b=%f, got %f", expectedB, vars["b"])
	}
	
	if vars["c"] != expectedC {
		t.Errorf("expected c=%f, got %f", expectedC, vars["c"])
	}
	
	if vars["x"] != expectedX {
		t.Errorf("expected x=%f, got %f", expectedX, vars["x"])
	}
}

func TestUnaryOperators(t *testing.T) {
	text := `BEGIN
	x := -5;
	y := +10;
	z := -x
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if vars["x"] != -5.0 {
		t.Errorf("expected x=-5, got %f", vars["x"])
	}
	
	if vars["y"] != 10.0 {
		t.Errorf("expected y=10, got %f", vars["y"])
	}
	
	if vars["z"] != 5.0 {
		t.Errorf("expected z=5, got %f", vars["z"])
	}
}

func TestComplexExpression(t *testing.T) {
	text := `BEGIN
	result := 2 + 3 * 4 - 10 / 2
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	expected := 9.0
	
	if vars["result"] != expected {
		t.Errorf("expected result=%f, got %f", expected, vars["result"])
	}
}

func TestVariableReuse(t *testing.T) {
	text := `BEGIN
	x := 5;
	x := x + 1;
	x := x * 2
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	expected := 12.0
	
	if vars["x"] != expected {
		t.Errorf("expected x=%f, got %f", expected, vars["x"])
	}
}

func TestMultipleVariables(t *testing.T) {
	text := `BEGIN
	a := 1;
	b := 2;
	c := 3;
	sum := a + b + c
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if vars["a"] != 1.0 || vars["b"] != 2.0 || vars["c"] != 3.0 || vars["sum"] != 6.0 {
		t.Errorf("unexpected variable values: %v", vars)
	}
}

func TestSingleStatement(t *testing.T) {
	text := `BEGIN
	x := 42
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if vars["x"] != 42.0 {
		t.Errorf("expected x=42, got %f", vars["x"])
	}
}

func TestEmptyStatements(t *testing.T) {
	text := `BEGIN
	;
	x := 5;
	;
END.`
	
	vars, err := Interpret(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if vars["x"] != 5.0 {
		t.Errorf("expected x=5, got %f", vars["x"])
	}
}

func TestDivisionByZeroError(t *testing.T) {
	text := `BEGIN
	x := 1 / 0
END.`
	
	_, err := Interpret(text)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

func TestUndefinedVariableError(t *testing.T) {
	text := `BEGIN
	x := y + 1
END.`
	
	_, err := Interpret(text)
	if err == nil {
		t.Error("expected undefined variable error")
	}
}

func TestSyntaxError(t *testing.T) {
	text := `BEGIN
	x := 
END.`
	
	_, err := Interpret(text)
	if err == nil {
		t.Error("expected syntax error")
	}
}

func TestMissingDot(t *testing.T) {
	text := `BEGIN
	x := 5
END`
	
	_, err := Interpret(text)
	if err == nil {
		t.Error("expected error for missing dot (.)")
	}
}