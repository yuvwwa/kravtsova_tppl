from .token import TokenType, Token

class Lexer():

    def __init__(self):
        self._text = ''
        self._pos = 0
        self._current_char = None

    def __number(self) -> str:
        result = []
        while (self._current_char is not None 
               and self._current_char.isdigit()):
            result.append(self._current_char)
            self.__forward()
        return ''.join(result)

    def next_token(self) -> Token:
        while self._current_char:
            if self._current_char.isspace(): 
                self.__skip()
                continue

            current_char = self._current_char
            if self._current_char.isdigit():
                return Token(TokenType.NUMBER, self.__number())
            if self._current_char in ['+', '-', '*', '/']:
                self.__forward()
                return Token(TokenType.OPERTATOR, current_char)
            if self._current_char == '(':
                self.__forward()
                return Token(TokenType.LPAREN, current_char)
            if self._current_char == ')':
                self.__forward()
                return Token(TokenType.RPAREN, current_char)
            raise SyntaxError()
        return Token(TokenType.EOL, "")
        
    def __skip(self):
        while (self._current_char is not None
               and self._current_char.isspace()):
            self.__forward()
    
    def __forward(self):
        self._pos += 1
        if self._pos >= len(self._text):
            self._current_char = None
        else:
            self._current_char = self._text[self._pos]
    
    def set_text(self, expr: str):
        self._text = expr
        self._pos = 0
        self._current_char = self._text[self._pos]