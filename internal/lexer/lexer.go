package lexer

type TokenType int

const (
	// Special tokens
	ILLEGAL TokenType = iota
	EOF

	// Identifiers and literals
	IDENT  // variable names
	STRING // 'hello world'
	INT    // 123

	// Keywords
	ENTRY    // Entry
	PRINT    // Print
	RETURN   // Return
	INT_TYPE // Int

	// Delimiters
	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }

	// Operators
	ASSIGN // =

	// Comments (we'll skip these in parsing)
	COMMENT
)

var keywords = map[string]TokenType{
	"Entry":  ENTRY,
	"Print":  PRINT,
	"Return": RETURN,
	"Int":    INT_TYPE,
}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII NUL character represents "EOF"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = Token{Type: ASSIGN, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '(':
		tok = Token{Type: LPAREN, Literal: string(l.ch), Line: l.line, Column: l.column}
	case ')':
		tok = Token{Type: RPAREN, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '{':
		tok = Token{Type: LBRACE, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '}':
		tok = Token{Type: RBRACE, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '\'':
		tok.Type = STRING
		tok.Literal = l.readString()
		tok.Line = l.line
		tok.Column = l.column
		l.readChar() // Skip the closing quote
		return tok
	case '/':
		if l.peekChar() == '/' {
			l.skipLineComment()
			return l.NextToken() // Skip comment and get next token
		} else if l.peekChar() == '*' {
			l.skipBlockComment()
			return l.NextToken() // Skip comment and get next token
		}
		tok = Token{Type: ILLEGAL, Literal: string(l.ch), Line: l.line, Column: l.column}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
		tok.Line = l.line
		tok.Column = l.column
	default:
		if isLetter(l.ch) {
			tok.Line = l.line
			tok.Column = l.column
			tok.Literal = l.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			tok.Line = l.line
			tok.Column = l.column
			return tok
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch), Line: l.line, Column: l.column}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1 // skip opening quote
	for {
		l.readChar()
		if l.ch == '\'' || l.ch == 0 {
			break
		}
		// Handle basic escape sequences
		if l.ch == '\\' && l.peekChar() != 0 {
			l.readChar() // Skip the escaped character
		}
	}
	str := l.input[position:l.position]
	return str
}

func (l *Lexer) readLineComment() string {
	position := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readBlockComment() string {
	position := l.position
	l.readChar() // skip '/'
	l.readChar() // skip '*'

	for {
		if l.ch == 0 {
			break
		}
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar() // skip '*'
			l.readChar() // skip '/'
			break
		}
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipBlockComment() {
	l.readChar() // skip '/'
	l.readChar() // skip '*'

	for {
		if l.ch == 0 {
			break
		}
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar() // skip '*'
			l.readChar() // skip '/'
			break
		}
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case STRING:
		return "STRING"
	case INT:
		return "INT"
	case ENTRY:
		return "ENTRY"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case INT_TYPE:
		return "INT_TYPE"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case ASSIGN:
		return "ASSIGN"
	case COMMENT:
		return "COMMENT"
	default:
		return "UNKNOWN"
	}
}
