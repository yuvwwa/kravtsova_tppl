package pascal

import "fmt"

type Parser struct {
	lexer        *Lexer
	currentToken Token
}

func NewParser(lexer *Lexer) (*Parser, error) {
	p := &Parser{lexer: lexer}
	token, err := lexer.NextToken()
	if err != nil {
		return nil, err
	}
	p.currentToken = token
	return p, nil
}

func (p *Parser) eat(tokenType TokenType) error {
	if p.currentToken.Type == tokenType {
		token, err := p.lexer.NextToken()
		if err != nil {
			return err
		}
		p.currentToken = token
		return nil
	}
	return fmt.Errorf("expected token type %d, got %d", tokenType, p.currentToken.Type)
}

func (p *Parser) Program() (Node, error) {
	node, err := p.compoundStatement()
	if err != nil {
		return nil, err
	}
	if err := p.eat(DOT); err != nil {
		return nil, err
	}
	return node, nil
}

func (p *Parser) compoundStatement() (Node, error) {
	if err := p.eat(BEGIN); err != nil {
		return nil, err
	}
	nodes, err := p.statementList()
	if err != nil {
		return nil, err
	}
	if err := p.eat(END); err != nil {
		return nil, err
	}
	
	root := Compound{Children: nodes}
	return root, nil
}

func (p *Parser) statementList() ([]Node, error) {
	node, err := p.statement()
	if err != nil {
		return nil, err
	}
	
	results := []Node{node}
	
	for p.currentToken.Type == SEMI {
		if err := p.eat(SEMI); err != nil {
			return nil, err
		}
		node, err := p.statement()
		if err != nil {
			return nil, err
		}
		results = append(results, node)
	}
	
	return results, nil
}

func (p *Parser) statement() (Node, error) {
	var node Node
	var err error
	
	if p.currentToken.Type == BEGIN {
		node, err = p.compoundStatement()
	} else if p.currentToken.Type == ID {
		node, err = p.assignmentStatement()
	} else {
		node = p.empty()
	}
	
	return node, err
}

func (p *Parser) assignmentStatement() (Node, error) {
	left, err := p.variable()
	if err != nil {
		return nil, err
	}
	token := p.currentToken
	if err := p.eat(ASSIGN); err != nil {
		return nil, err
	}
	right, err := p.expr()
	if err != nil {
		return nil, err
	}
	return Assign{Left: left, Op: token, Right: right}, nil
}

func (p *Parser) variable() (Var, error) {
	node := Var{Token: p.currentToken, Value: p.currentToken.Value}
	if err := p.eat(ID); err != nil {
		return Var{}, err
	}
	return node, nil
}

func (p *Parser) empty() Node {
	return NoOp{}
}

func (p *Parser) expr() (Node, error) {
	node, err := p.term()
	if err != nil {
		return nil, err
	}
	
	for p.currentToken.Type == PLUS || p.currentToken.Type == MINUS {
		token := p.currentToken
		if token.Type == PLUS {
			if err := p.eat(PLUS); err != nil {
				return nil, err
			}
		} else {
			if err := p.eat(MINUS); err != nil {
				return nil, err
			}
		}
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		node = BinOp{Left: node, Op: token, Right: right}
	}
	
	return node, nil
}

func (p *Parser) term() (Node, error) {
	node, err := p.factor()
	if err != nil {
		return nil, err
	}
	
	for p.currentToken.Type == MUL || p.currentToken.Type == DIV {
		token := p.currentToken
		if token.Type == MUL {
			if err := p.eat(MUL); err != nil {
				return nil, err
			}
		} else {
			if err := p.eat(DIV); err != nil {
				return nil, err
			}
		}
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		node = BinOp{Left: node, Op: token, Right: right}
	}
	
	return node, nil
}

func (p *Parser) factor() (Node, error) {
	token := p.currentToken
	
	if token.Type == PLUS {
		if err := p.eat(PLUS); err != nil {
			return nil, err
		}
		node, err := p.factor()
		if err != nil {
			return nil, err
		}
		return UnaryOp{Op: token, Expr: node}, nil
	} else if token.Type == MINUS {
		if err := p.eat(MINUS); err != nil {
			return nil, err
		}
		node, err := p.factor()
		if err != nil {
			return nil, err
		}
		return UnaryOp{Op: token, Expr: node}, nil
	} else if token.Type == INTEGER {
		if err := p.eat(INTEGER); err != nil {
			return nil, err
		}
		return Number{Token: token, Value: token.Value}, nil
	} else if token.Type == LPAREN {
		if err := p.eat(LPAREN); err != nil {
			return nil, err
		}
		node, err := p.expr()
		if err != nil {
			return nil, err
		}
		if err := p.eat(RPAREN); err != nil {
			return nil, err
		}
		return node, nil
	} else {
		node, err := p.variable()
		if err != nil {
			return nil, err
		}
		return node, nil
	}
}

func (p *Parser) Parse() (Node, error) {
	node, err := p.Program()
	if err != nil {
		return nil, err
	}
	if p.currentToken.Type != EOF {
		return nil, fmt.Errorf("unexpected token after program end")
	}
	return node, nil
}