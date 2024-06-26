package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
	Children() []Node
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) Children() []Node {

	nodes := make([]Node, len(p.Statements))
	for i, s := range p.Statements {
		nodes[i] = s
	}

	return nodes
}

// let name = value;
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (p *LetStatement) Children() []Node {
	return []Node{p.Value}
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (p *Identifier) Children() []Node {
	return []Node{}
}
func (i *Identifier) String() string       { return i.Value }
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// return ReturnValue;
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (p *ReturnStatement) Children() []Node {
	return []Node{p.ReturnValue}
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (p *ExpressionStatement) Children() []Node {
	return []Node{p.Expression}
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

type IntegerLiteral struct {
	Token token.Token // the token.INT token
	Value int64
}

func (p *IntegerLiteral) Children() []Node {
	return []Node{}
}
func (i *IntegerLiteral) String() string       { return i.Token.Literal }
func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }

type StringLiteral struct {
	Token token.Token // the token.STRING token
	Value string
}

func (p *StringLiteral) Children() []Node {
	return []Node{}
}
func (i *StringLiteral) String() string       { return i.Token.Literal }
func (i *StringLiteral) expressionNode()      {}
func (i *StringLiteral) TokenLiteral() string { return i.Token.Literal }

type PrefixExpression struct {
	Token    token.Token // the PREFIX token, e.g. !
	Operator string
	Right    Expression
}

func (p *PrefixExpression) Children() []Node {
	return []Node{p.Right}
}
func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // the Infix token, e.g. +, *, /, etc
	Left     Expression
	Operator string
	Right    Expression
}

func (p *InfixExpression) Children() []Node {
	return []Node{p.Left, p.Right}
}
func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (p *Boolean) Children() []Node {
	return []Node{}
}
func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (p *IfExpression) Children() []Node {
	return []Node{p.Condition, p.Consequence, p.Alternative}
}
func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (p *BlockStatement) Children() []Node {
	nodes := make([]Node, len(p.Statements))
	for i, n := range p.Statements {
		nodes[i] = n
	}

	return nodes
}
func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (p *FunctionLiteral) Children() []Node {
	return []Node{p.Body}
}
func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (p *CallExpression) Children() []Node {
	nodes := make([]Node, len(p.Arguments)+1)
	nodes[0] = p.Function
	for i, n := range p.Arguments {
		nodes[i+1] = n
	}

	return nodes
}
func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (p *ArrayLiteral) Children() []Node {
	nodes := make([]Node, len(p.Elements))
	for i, n := range p.Elements {
		nodes[i] = n
	}

	return nodes
}
func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type IndexExpression struct {
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

func (p *IndexExpression) Children() []Node {
	return []Node{p.Left, p.Index}
}
func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (p *HashLiteral) Children() []Node {

	nodes := []Node{}
	for k, v := range p.Pairs {
		nodes = append(nodes, k, v)
	}
	return nodes
}
func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
