package lexer

import "github.com/adoringonion/golang_monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = NewToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = NewToken(token.SEMICOLON, l.ch)
	case ':':
		tok = NewToken(token.COLON, l.ch)
	case '(':
		tok = NewToken(token.LPAREN, l.ch)
	case ')':
		tok = NewToken(token.RPAREN, l.ch)
	case '[':
		tok = NewToken(token.LBRACKET, l.ch)
	case ']':
		tok = NewToken(token.RBRACKET, l.ch)
	case ',':
		tok = NewToken(token.COMMA, l.ch)
	case '+':
		tok = NewToken(token.PLUS, l.ch)
	case '-':
		tok = NewToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = NewToken(token.BANG, l.ch)
		}
	case '/':
		tok = NewToken(token.SLASH, l.ch)
	case '*':
		tok = NewToken(token.ASTERISK, l.ch)
	case '<':
		tok = NewToken(token.LT, l.ch)
	case '>':
		tok = NewToken(token.RT, l.ch)
	case '{':
		tok = NewToken(token.LBRACE, l.ch)
	case '}':
		tok = NewToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	positon := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[positon:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}

}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

var keywards = map[string]token.TokenType{
	"fn":     token.FUNCTION,
	"let":    token.LET,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"if":     token.IF,
	"else":   token.ELSE,
	"return": token.RETURN,
}

func LookupIdent(ident string) token.TokenType {
	if tok, ok := keywards[ident]; ok {
		return tok
	}

	return token.IDENT
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
