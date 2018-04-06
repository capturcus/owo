package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
		`|(^[#;].*$)` +
		`|(?P<Ident>[a-zA-Z][a-zA-Z_\d]*)` +
		`|(?P<String>"(?:\\.|[^"])*")` +
		`|(?P<Float>\d+(?:\.\d+)?)` +
		`|(?P<Punct>[][=])`,
)))

type INI struct {
	Properties []*Property `{ @@ }`
	Sections   []*Section  `{ @@ Newline }`
}

type Section struct {
	Identifier string      `"[" @Ident "]" Newline`
	Properties []*Property `Indent { @@ Newline } Dedent`
}

type Property struct {
	Key   string `@Ident "="`
	Value *Value `@@`
}

type Value struct {
	String *string  `  @String`
	Number *float64 `| @Float`
}

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
			fmt.Println("INDENT")
			numTabs = lineTabs
			line = string(INDENT) + line
		}
		if lineTabs < numTabs {
			fmt.Println("DEDENT")
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
	parser, err := participle.Build(&INI{}, iniLexer)
	if err != nil {
		panic(err)
	}
	ini := &INI{}
	err = parser.Parse(strings.NewReader(dentedSource), ini)
	if err != nil {
		panic("PARSING ERROR" + err.Error())
	}
	repr.Println(ini, repr.Indent("  "), repr.OmitEmpty(true))
}
