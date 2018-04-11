package main

import (
	"io/ioutil"
	"os"
	"owo/gocc/lexer"
	"owo/gocc/parser"
	"strings"

	"github.com/alecthomas/repr"
)

func Detonate(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	INDENT = rune(17)
	DEDENT = rune(16)
)

func countTabs(line string) int {
	for i, char := range line {
		if char != '\t' {
			return i
		}
	}
	return 0
}

func dentSource(source string) string {
	lines := strings.Split(source, "\n")
	numTabs := 0
	dentedLines := make([]string, 0)
	for _, line := range lines {
		lineTabs := countTabs(line)
		if lineTabs > numTabs {
			line = strings.Repeat(string(INDENT), lineTabs-numTabs) + line
			numTabs = lineTabs
		}
		if lineTabs < numTabs {
			line = strings.Repeat(string(DEDENT)+"\n", numTabs-lineTabs) + line
			numTabs = lineTabs
		}
		dentedLines = append(dentedLines, line)
	}
	return strings.Join(dentedLines, "\n")
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	Detonate(err)
	input = []byte(dentSource(string(input)))
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	st, err := p.Parse(lex)
	if err != nil {
		panic("PARSE ERROR: " + err.Error())
	}
	repr.Println("OK", st, repr.Indent("  "))
}
