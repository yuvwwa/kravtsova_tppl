package pascal

type Node interface {
	node()
}

type Number struct {
	Token Token
	Value string
}

func (n Number) node() {}

type BinOp struct {
	Left  Node
	Op    Token
	Right Node
}

func (b BinOp) node() {}

type UnaryOp struct {
	Op   Token
	Expr Node
}

func (u UnaryOp) node() {}

type Compound struct {
	Children []Node
}

func (c Compound) node() {}

type Assign struct {
	Left  Var
	Op    Token
	Right Node
}

func (a Assign) node() {}

type Var struct {
	Token Token
	Value string
}

func (v Var) node() {}

type NoOp struct{}

func (n NoOp) node() {}