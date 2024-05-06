package tokenizer

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	ATOM TokenType = iota
	LPAREN
	RPAREN
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	Input string
	pos   int
}

func (t *Tokenizer) next() rune {
	if t.pos >= len(t.Input) {
		return 0
	}
	return rune(t.Input[t.pos])
}

func (t *Tokenizer) consume() {
	t.pos++
}

func (t *Tokenizer) skipWhitespace() {
	for unicode.IsSpace(t.next()) {
		t.consume()
	}
}

func (t *Tokenizer) Tokenize() ([]Token, error) {
	var tokens []Token
	parens := 0
	for t.pos < len(t.Input) {
		t.skipWhitespace()
		switch ch := t.next(); ch {
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: "("})
			parens++
			t.consume()
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
			parens--
			t.consume()
		default:
			if !(unicode.IsSpace(ch) || ch == '\n') {
				var sb strings.Builder
				for !unicode.IsSpace(t.next()) && t.next() != '(' && t.next() != ')' {
					sb.WriteRune(t.next())
					t.consume()
				}
				tokens = append(tokens, Token{Type: ATOM, Value: sb.String()})
			} else {
				t.consume()
			}
		}
	}
	if parens != 0 {
		return nil, fmt.Errorf("unbalanced parentheses")
	}
	return tokens, nil
}
