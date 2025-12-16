package pascal

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
	parser    *Parser
	variables map[string]float64
}

func NewInterpreter(parser *Parser) *Interpreter {
	return &Interpreter{
		parser:    parser,
		variables: make(map[string]float64),
	}
}

func (i *Interpreter) visit(node Node) (float64, error) {
	switch n := node.(type) {
	case Number:
		return i.visitNumber(n)
	case BinOp:
		return i.visitBinOp(n)
	case UnaryOp:
		return i.visitUnaryOp(n)
	case Compound:
		return i.visitCompound(n)
	case Assign:
		return i.visitAssign(n)
	case Var:
		return i.visitVar(n)
	case NoOp:
		return i.visitNoOp(n)
	default:
		return 0, fmt.Errorf("unknown node type")
	}
}

func (i *Interpreter) visitNumber(node Number) (float64, error) {
	value, err := strconv.ParseFloat(node.Value, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (i *Interpreter) visitBinOp(node BinOp) (float64, error) {
	left, err := i.visit(node.Left)
	if err != nil {
		return 0, err
	}
	right, err := i.visit(node.Right)
	if err != nil {
		return 0, err
	}
	
	switch node.Op.Type {
	case PLUS:
		return left + right, nil
	case MINUS:
		return left - right, nil
	case MUL:
		return left * right, nil
	case DIV:
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown operator")
	}
}

func (i *Interpreter) visitUnaryOp(node UnaryOp) (float64, error) {
	value, err := i.visit(node.Expr)
	if err != nil {
		return 0, err
	}
	
	switch node.Op.Type {
	case PLUS:
		return value, nil
	case MINUS:
		return -value, nil
	default:
		return 0, fmt.Errorf("unknown unary operator")
	}
}

func (i *Interpreter) visitCompound(node Compound) (float64, error) {
	for _, child := range node.Children {
		_, err := i.visit(child)
		if err != nil {
			return 0, err
		}
	}
	return 0, nil
}

func (i *Interpreter) visitAssign(node Assign) (float64, error) {
	varName := node.Left.Value
	value, err := i.visit(node.Right)
	if err != nil {
		return 0, err
	}
	i.variables[varName] = value
	return value, nil
}

func (i *Interpreter) visitVar(node Var) (float64, error) {
	varName := node.Value
	value, ok := i.variables[varName]
	if !ok {
		return 0, fmt.Errorf("undefined variable: %s", varName)
	}
	return value, nil
}

func (i *Interpreter) visitNoOp(node NoOp) (float64, error) {
	return 0, nil
}

func (i *Interpreter) Interpret() (map[string]float64, error) {
	tree, err := i.parser.Parse()
	if err != nil {
		return nil, err
	}
	_, err = i.visit(tree)
	if err != nil {
		return nil, err
	}
	return i.variables, nil
}

func Interpret(text string) (map[string]float64, error) {
	lexer := NewLexer(text)
	parser, err := NewParser(lexer)
	if err != nil {
		return nil, err
	}
	interpreter := NewInterpreter(parser)
	return interpreter.Interpret()
}