lisp-lang
---------

This project is a simple Lisp interpreter written in Go. It supports basic Lisp features such as defining and calling functions, and arithmetic operations.
This a purse hobby project. May add more features in the future.

## Features

- Define and call functions using the `defun` keyword.
- Perform arithmetic operations such as addition and multiplication.
- Parse and evaluate multiple statements.

## Usage

To use the interpreter, you need to write your Lisp code as a string and pass it to the `Tokenizer`:

```go
t := tokenizer.Tokenizer{Input: "(defun test (* 2 (+ 3 2))) (test)"}
```

Then, you can tokenize the input:

```go
tokens, err := t.Tokenize()
if err != nil {
    fmt.Println(err)
    return
}
```

After tokenizing the input, you can parse it:

```go
p := parser.Parser{Tokens: tokens}
asts := p.Parse()
```

Finally, you can create an interpreter and visit each abstract syntax tree (AST):

```go
v := interpreter.NewInterpreter()
for _, ast := range asts {
    result := v.Visit(ast)
    fmt.Println(result.Display())
}
```

## Example

Here is an example of defining a function and calling it:

```lisp
(defun test (* 2 (+ 3 2))) (test)
```

This will define a function named `test` that multiplies 2 by the sum of 3 and 2. Then, it calls the `test` function.

## Installation

To install the project, you can clone the repository and build it using Go:

```bash
git clone https://github.com/allen13/lisp-lang.git
cd lisp-lang
go build
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
