package parser

import "lisp-lang/tokenizer"

type NodeType int

const (
	ATOM NodeType = iota
	LIST
	FUNCTION
)

type Node struct {
	Type     NodeType
	Value    string
	Children []*Node
}

type Parser struct {
	Tokens []tokenizer.Token
	pos    int
}

func (p *Parser) next() tokenizer.Token {
	if p.pos >= len(p.Tokens) {
		return tokenizer.Token{}
	}
	return p.Tokens[p.pos]
}

func (p *Parser) consume() tokenizer.Token {
	token := p.next()
	p.pos++
	return token
}

func (p *Parser) parseAtom() *Node {
	token := p.consume()
	return &Node{Type: ATOM, Value: token.Value}
}

func (p *Parser) parseList() *Node {
	p.consume() // consume '('
	var children []*Node
	for p.next().Type != tokenizer.RPAREN {
		children = append(children, p.parseExpr())
	}
	p.consume() // consume ')'

	// Check if this is a function definition
	if len(children) > 0 && children[0].Value == "defun" {
		return &Node{Type: FUNCTION, Children: children[1:]}
	}

	return &Node{Type: LIST, Children: children}
}

func (p *Parser) parseExpr() *Node {
	switch p.next().Type {
	case tokenizer.LPAREN:
		return p.parseList()
	default:
		return p.parseAtom()
	}
}

func (p *Parser) Parse() []*Node {
	var nodes []*Node
	for p.pos < len(p.Tokens) {
		nodes = append(nodes, p.parseExpr())
	}
	return nodes
}
