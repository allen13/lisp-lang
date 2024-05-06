package interpreter

import (
	"fmt"
	"lisp-lang/parser"
	"strconv"
)

type Interpreter struct {
	node      *parser.Node
	Functions map[string]*parser.Node
}

type Result struct {
	Type        string
	IntValue    int
	StringValue string
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		Functions: make(map[string]*parser.Node),
	}
}

func (r Result) Display() string {
	switch r.Type {
	case "int":
		return strconv.Itoa(r.IntValue)
	case "string":
		return r.StringValue
	default:
		return fmt.Sprintf("%s: %s", r.Type, r.StringValue)
	}
}

func (v *Interpreter) visitAtom(node *parser.Node) Result {
	value, err := strconv.Atoi(node.Value)
	if err != nil {
		// If the value is not a number, return it as a string
		return Result{Type: "string", StringValue: node.Value}
	}
	return Result{Type: "int", IntValue: value}
}

func (v *Interpreter) visitList(node *parser.Node) Result {
	head := node.Children[0]
	tail := node.Children[1:]
	switch head.Value {
	case "+":
		// Calculate the sum of the values of the tail of the list
		sum := 0
		for _, child := range tail {
			result := v.Visit(child)
			if result.Type != "int" {
				return Result{Type: "error", StringValue: "invalid argument to '+'"}
			}
			sum += result.IntValue
		}
		return Result{Type: "int", IntValue: sum}
	case "*":
		// Calculate the product of the values of the tail of the list
		product := 1
		for _, child := range tail {
			result := v.Visit(child)
			if result.Type != "int" {
				return Result{Type: "error", StringValue: "invalid argument to '*'"}
			}
			product *= result.IntValue
		}
		return Result{Type: "int", IntValue: product}
	default:
		// Treat this as a function call
		return v.visitFunction(node)
	}
}

func (v *Interpreter) Visit(node *parser.Node) Result {
	switch node.Type {
	case parser.ATOM:
		return v.visitAtom(node)
	case parser.LIST:
		return v.visitList(node)
	case parser.FUNCTION:
		v.Functions[node.Children[0].Value] = node.Children[1]
		return Result{Type: "string", StringValue: fmt.Sprintf("function defined: %s", node.Children[0].Value)}
	default:
		fmt.Println(node.Type)
		return Result{Type: "error", StringValue: "unknown node type"}
	}
}

func (v *Interpreter) visitFunction(node *parser.Node) Result {
	// Check if the function is defined
	functionName := node.Children[0].Value
	functionNode, ok := v.Functions[functionName]
	if !ok {
		return Result{Type: "error", StringValue: "function not defined"}
	}

	return v.Visit(functionNode)
}
