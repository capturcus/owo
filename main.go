package main

import (
	"io/ioutil"
	"os"
	"owo/grammar"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/repr"
)

const (
	INDENT = rune(17)
	DEDENT = rune(16)
)

// A custom lexer for INI files. This illustrates a relatively complex Regexp lexer, as well
// as use of the Unquote filter, which unquotes string tokens.
var iniLexer = lexer.Unquote(lexer.Must(lexer.Regexp(
	`(?m)` +
		`(\t+)` +
		`|(?P<Indent>\x11)` +
		`|(?P<Dedent>\x10)` +
		`|(?P<Newline>\n)` +
		`|(?P<Colon>:)` +
		`|(?P<Comma>,)` +
		`|(?P<Spaces>[ ]+)` +
		`|(^[#;].*$)` +
		`|(?P<Ident>[a-zA-Z][a-zA-Z_\d]*)` +
		`|(?P<String>"(?:\\.|[^"])*")` +
		`|(?P<Float>\d+(?:\.\d+)?)` +
		`|(?P<Punct>[][=])`,
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
			numTabs = lineTabs
			line = string(INDENT) + line
		}
		if lineTabs < numTabs {
			numTabs = lineTabs
			line = line + string(DEDENT)
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
