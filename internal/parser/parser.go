package parser

import (
	"dreadlang/internal/lexer"
	"fmt"
)

// AST Node types
type Node interface {
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += s.String()
	}
	return out
}

// Parameter represents a function parameter
type Parameter struct {
	Name string
	Type string
}

func (p *Parameter) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Type)
}

// Statements
type FunctionStatement struct {
	IsEntry    bool
	Name       string
	Parameters []*Parameter
	ReturnType string
	Body       *BlockStatement
}

func (fs *FunctionStatement) statementNode() {}
func (fs *FunctionStatement) String() string {
	var keyword string
	if fs.IsEntry {
		keyword = "Entry"
	} else {
		keyword = "Function"
	}

	var params string
	for i, param := range fs.Parameters {
		if i > 0 {
			params += ", "
		}
		params += param.String()
	}

	return fmt.Sprintf("%s %s(%s) (%s) %s", keyword, fs.Name, params, fs.ReturnType, fs.Body.String())
}

type BlockStatement struct {
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) String() string {
	var out string
	out += "{"
	for _, s := range bs.Statements {
		out += s.String()
	}
	out += "}"
	return out
}

type AssignStatement struct {
	Name  string
	Value Expression
}

func (as *AssignStatement) statementNode() {}
func (as *AssignStatement) String() string {
	return fmt.Sprintf("%s = %s", as.Name, as.Value.String())
}

type CallStatement struct {
	Function  string
	Arguments []Expression
}

func (cs *CallStatement) statementNode() {}
func (cs *CallStatement) String() string {
	var args string
	for i, arg := range cs.Arguments {
		if i > 0 {
			args += ", "
		}
		args += arg.String()
	}
	return fmt.Sprintf("%s(%s)", cs.Function, args)
}

// Expressions
type StringLiteral struct {
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) String() string {
	return fmt.Sprintf("'%s'", sl.Value)
}

type IntegerLiteral struct {
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) String() string {
	return i.Value
}

type CallExpression struct {
	Function  string
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) String() string {
	var args string
	for i, arg := range ce.Arguments {
		if i > 0 {
			args += ", "
		}
		args += arg.String()
	}
	return fmt.Sprintf("%s(%s)", ce.Function, args)
}

// Parser
type Parser struct {
	l *lexer.Lexer

	curToken  lexer.Token
	peekToken lexer.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != lexer.EOF {
		// Skip comments
		if p.curToken.Type == lexer.COMMENT {
			p.nextToken()
			continue
		}

		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case lexer.ENTRY:
		return p.parseFunctionStatement(true)
	case lexer.FUNCTION:
		return p.parseFunctionStatement(false)
	default:
		return p.parseBlockStatement()
	}
}

func (p *Parser) parseFunctionStatement(isEntry bool) Statement {
	stmt := &FunctionStatement{
		IsEntry: isEntry,
	}

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}

	stmt.Name = p.curToken.Literal

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	// Parse parameters
	stmt.Parameters = p.parseParameters()

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	// Handle return type - three possible syntaxes:
	// 1. () (Type)  - parenthesized return type
	// 2. () Type    - bare return type
	// 3. () {       - no return type (defaults to Void)
	if p.peekToken.Type == lexer.LPAREN {
		// Syntax: () (Type)
		p.nextToken() // consume LPAREN
		if !p.expectPeek(lexer.INT_TYPE) && !p.expectPeek(lexer.STRING_TYPE) && !p.expectPeek(lexer.VOID_TYPE) {
			return nil
		}
		stmt.ReturnType = p.curToken.Literal
		if !p.expectPeek(lexer.RPAREN) {
			return nil
		}
	} else if p.peekToken.Type == lexer.INT_TYPE || p.peekToken.Type == lexer.STRING_TYPE || p.peekToken.Type == lexer.VOID_TYPE {
		// Syntax: () Type
		p.nextToken()
		stmt.ReturnType = p.curToken.Literal
	} else {
		// No return type specified, default to Void
		stmt.ReturnType = "Void"
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseParameters() []*Parameter {
	parameters := []*Parameter{}

	// If the next token is RPAREN, there are no parameters
	if p.peekToken.Type == lexer.RPAREN {
		return parameters
	}

	// Move to the first parameter
	p.nextToken()

	// Parse first parameter
	param := p.parseParameter()
	if param != nil {
		parameters = append(parameters, param)
	}

	// Parse remaining parameters
	for p.peekToken.Type == lexer.COMMA {
		p.nextToken() // consume the comma
		p.nextToken() // move to next parameter
		param := p.parseParameter()
		if param != nil {
			parameters = append(parameters, param)
		}
	}

	return parameters
}

func (p *Parser) parseParameter() *Parameter {
	// Support syntax: Type name (e.g., "String input_str")
	if p.curToken.Type == lexer.STRING_TYPE || p.curToken.Type == lexer.INT_TYPE {
		param := &Parameter{
			Type: p.curToken.Literal,
		}

		if !p.expectPeek(lexer.IDENT) {
			return nil
		}

		param.Name = p.curToken.Literal
		return param
	}

	// Support syntax: name Type (e.g., "input_str String")
	if p.curToken.Type == lexer.IDENT {
		param := &Parameter{
			Name: p.curToken.Literal,
		}

		if !p.expectPeek(lexer.STRING_TYPE) && !p.expectPeek(lexer.INT_TYPE) {
			return nil
		}

		param.Type = p.curToken.Literal
		return param
	}

	return nil
}

func (p *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{}
	block.Statements = []Statement{}

	p.nextToken()

	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		// Skip comments
		if p.curToken.Type == lexer.COMMENT {
			p.nextToken()
			continue
		}

		stmt := p.parseInnerStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseInnerStatement() Statement {
	switch p.curToken.Type {
	case lexer.IDENT:
		if p.peekToken.Type == lexer.ASSIGN {
			return p.parseAssignStatement()
		} else if p.peekToken.Type == lexer.LPAREN {
			// This is a function call statement
			return p.parseCallStatement()
		}
		return nil
	case lexer.PRINT, lexer.RETURN:
		return p.parseCallStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAssignStatement() Statement {
	stmt := &AssignStatement{}
	stmt.Name = p.curToken.Literal

	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression()

	return stmt
}

func (p *Parser) parseCallStatement() Statement {
	stmt := &CallStatement{}
	stmt.Function = p.curToken.Literal

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	stmt.Arguments = p.parseArgumentList()

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	return stmt
}

func (p *Parser) parseArgumentList() []Expression {
	args := []Expression{}

	// If the next token is RPAREN, there are no arguments
	if p.peekToken.Type == lexer.RPAREN {
		return args
	}

	// Move to the first argument
	p.nextToken()

	// Parse first argument
	arg := p.parseExpression()
	if arg != nil {
		args = append(args, arg)
	}

	// Parse remaining arguments
	for p.peekToken.Type == lexer.COMMA {
		p.nextToken() // consume the comma
		p.nextToken() // move to next argument
		arg := p.parseExpression()
		if arg != nil {
			args = append(args, arg)
		}
	}

	return args
}

func (p *Parser) parseExpression() Expression {
	switch p.curToken.Type {
	case lexer.STRING:
		return &StringLiteral{Value: p.curToken.Literal}
	case lexer.INT:
		// For MVP, we'll just store as string and handle conversion later
		return &StringLiteral{Value: p.curToken.Literal}
	case lexer.IDENT:
		// Check if this is a function call
		if p.peekToken.Type == lexer.LPAREN {
			return p.parseCallExpression()
		}
		return &Identifier{Value: p.curToken.Literal}
	default:
		return nil
	}
}

func (p *Parser) parseCallExpression() Expression {
	expr := &CallExpression{}
	expr.Function = p.curToken.Literal

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	expr.Arguments = p.parseArgumentList()

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	return expr
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
