// pascal/token.go
package pascal

type TokenType int

const (
	INTEGER TokenType = iota
	PLUS
	MINUS
	MUL
	DIV
	LPAREN
	RPAREN
	ID
	ASSIGN
	BEGIN
	END
	SEMI
	DOT
	EOF
)

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(tokenType TokenType, value string) Token {
	return Token{
		Type:  tokenType,
		Value: value,
	}
}

func (t Token) String() string {
	return "Token(" + t.Value + ")"
}