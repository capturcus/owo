package main

import (
	"io/ioutil"
	"os"
	"owo/gocc/lexer"
	"owo/gocc/parser"

	"github.com/alecthomas/repr"
)

func Detonate(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	Detonate(err)
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	st, err := p.Parse(lex)
	if err != nil {
		panic(err)
	}
	repr.Println("OK", st, repr.Indent("  "))
}
