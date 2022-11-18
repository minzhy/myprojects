package ast

import "myprojects/token"

// 每个ast树的节点都要继承，因为每个节点至少一个Token，所以一定有TokenLiteral
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token // token.LET 词法单元，所以这个Token一定是Let
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type Expression interface {
	Node
	expressionNode()
}

// Identifier 其实算作表达式的其中一种！毕竟也可以当成是表达式。expression并不一定是Identify
type Identifier struct {
	Token token.Token // token.IDENT 词法单元
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// 整个程序是有一堆的statement组成的！！program相当于是一个根节点啊
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
