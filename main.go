package main

import (
	"fmt"
	"lisp-lang/interpreter"
	"lisp-lang/tokenizer"
)
import "lisp-lang/parser"

func printNode(node *parser.Node, indent string) {
	switch node.Type {
	case parser.ATOM:
		fmt.Println(indent, "ATOM:", node.Value)
	case parser.LIST:
		fmt.Println(indent, "LIST:")
		for _, child := range node.Children {
			printNode(child, indent+"  ")
		}
	case parser.FUNCTION:
		fmt.Println(indent, "FUNCTION:", node.Value)
		for _, child := range node.Children {
			printNode(child, indent+"  ")
		}
	}
}

func main() {
	// Assume t and p are the tokenizer and parser from the previous code
	t := tokenizer.Tokenizer{Input: "(defun test (* 2 (+ 3 2))) (test)"}
	tokens, err := t.Tokenize()
	if err != nil {
		fmt.Println(err)
		return
	}
	p := parser.Parser{Tokens: tokens}
	asts := p.Parse()
	v := interpreter.NewInterpreter()
	for _, ast := range asts {
		printNode(ast, "")
		result := v.Visit(ast)
		fmt.Println(result.Display())
	}
}
