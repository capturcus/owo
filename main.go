package main

import (
	"io/ioutil"
	"os"
	"owo/grammar"
	"strings"

	"owo/participle"
	"owo/participle/lexer"

	"owo/repr"
)

const (
	INDENT = rune(17)
	DEDENT = rune(16)
)

// A custom lexer for INI files. This illustrates a relatively complex Regexp lexer, as well
// as use of the Unquote filter, which unquotes string tokens.
var iniLexer = lexer.Unquote(lexer.Must(lexer.Regexp(
	grammar.LEXER_STRING,
)))

func countTabs(line string) int {
	for i, char := range line {
		if char != '\t' {
			return i
		}
	}
	return 0
}

func dentSource(source string) (string, error) {
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
	return strings.Join(dentedLines, "\n"), nil
}

func main() {
	bytez, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	strSource := string(bytez)
	dentedSource, err := dentSource(strSource)
	if err != nil {
		panic(err)
	}
	parser, err := participle.Build(&grammar.Program{}, iniLexer)
	if err != nil {
		panic(err)
	}
	ini := &grammar.Program{}
	err = parser.Parse(strings.NewReader(dentedSource), ini)
	if err != nil {
		panic("PARSING ERROR" + err.Error())
	}
	repr.Println(ini, repr.Indent("  "), repr.OmitEmpty(true))
}
