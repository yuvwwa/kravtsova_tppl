from .ast import Node, Number, BinOp
from .parser import Parser
from .lexer import Lexer

class NodeVisitor:

    def visit(self):
        pass

class Interpreter():

    def __init__(self):
        self._lexer = Lexer()   
        self._parser = Parser(self._lexer)

    def visit(self, node: Node):
        if isinstance(node, Number):
            return self.visit_number(node)
        if isinstance(node, BinOp):
            return self.visit_binop(node)

    def visit_number(self, node: Number):
        return float(node.token.value)
    
    def visit_binop(self, node: BinOp):
        match node.op.value:
            case '+':
                return self.visit(node.left) + self.visit(node.right)
            case '-':
                return self.visit(node.left) - self.visit(node.right)
            case '*':
                return self.visit(node.left) * self.visit(node.right)
            case '/':
                return self.visit(node.left) / self.visit(node.right)
            case _:
                raise ValueError('Invalid operator')

    def eval(self, text: str) -> float:
        tree = self._parser.parse(text)
        return self.visit(tree)