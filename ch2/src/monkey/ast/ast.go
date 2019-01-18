package ast

import (
	"bytes"
	"goInterpreter/ch2/src/monkey/token"
)

type Node interface{
	TokenLiteral() string // For debugging
	String() string
}

type Statement interface {
	// Include all the methods from the 'Node' interface
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Root node of AST
// Every program is a series of statements
// Statements is a slice of AST Nodes that implement
// the statement Interface
type Program struct{
	Statements []Statement
}

func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0{
		return p.Statements[0].TokenLiteral()
	} else{
		return ""
	}
}

func (p *Program) String() string{
	var out bytes.Buffer

	for _, s := range p.Statements{
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct{
	Token token.Token // token.LET
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

func(ls *LetStatement) String() string{
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil{
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode(){}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}

func (i *Identifier) String() string{
	return i.Value
}

type ReturnStatement struct{
	Token token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode(){}
func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Literal}
func (rs *ReturnStatement) String() string{
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil{
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}


type ExpresionStatement struct{
	Token token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpresionStatement) statementNode(){}
func (es *ExpresionStatement) TokenLiteral() string {return es.Token.Literal}

func (es *ExpresionStatement) String() string {
	if es.Expression != nil{
		return es.Expression.String()
	}
	return ""
}


type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {return il.Token.Literal}
func (il *IntegerLiteral) String() string {return il.Token.Literal}