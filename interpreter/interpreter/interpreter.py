from .token import TokenType, Token

class Interpreter():
    def __init__(self):
        self._pos = 0
        self._current_token = None
        self._current_char = None
        self._text = ""

    def __check_token_type(self, type_: TokenType):
        if self._current_token.type_ == type_:
            self._current_token = self.__next_token()
        else:
            raise SyntaxError("Invalid token type")
        
    def __forward(self):
        self._pos += 1
        if self._pos > len(self._text)-1:
            self._current_char = None
        else: 
            self._current_char = self._text[self._pos]

    def __skip(self):
        while self._current_char is not None and self._current_char.isspace():
            self.__forward()

    def __number(self) -> str:
        result = []
        while (self._current_char is not None and
               self._current_char.isdigit()):
            result.append(self._current_char)
            self.__forward()

        return "".join(result)

    def __next_token(self) -> Token:
        while self._current_char:
            if self._current_char.isspace():
                self.__skip()
                continue

            current_char = self._current_char

            if self._current_char.isdigit():
                return Token(TokenType.NUMBER, self.__number())
            
            if self._current_char in ["+", "-"]:
                self.__forward()
                return Token(TokenType.OPERATOR, current_char)
        
            raise SyntaxError()

    def __term(self) -> float:
        token = self._current_token
        self.__check_token_type(TokenType.NUMBER)

        return float(token.value)


    def __expr(self) -> float:
        self._current_token = self.__next_token()
        result = self.__term()

        while (self._current_token is not None and self._current_token.type_ == TokenType.OPERATOR):
            token = self._current_token
            if token.value == "+":
                self.__check_token_type(TokenType.OPERATOR)
                result += self.__term()
            elif token.value == "-":
                self.__check_token_type(TokenType.OPERATOR)
                result -= self.__term()

        return result


    def eval(self, expr:str) -> float:
        self._text = expr
        self._pos = 0
        self._current_char = self._text[self._pos]
        self._current_token = None
        return self.__expr()
    