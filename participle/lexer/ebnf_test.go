package lexer

import (
	"strings"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
)

func TestBuilder(t *testing.T) {
	tests := []struct {
		name      string
		grammar   string
		source    string
		tokens    []string
		roots     []string
		failBuild bool
		fail      bool
	}{
		{
			name:      "BadEBNF",
			grammar:   "Production = helper .",
			failBuild: true,
		},
		{
			name:    "EmptyProductionErrorsWithInput",
			grammar: `Extra = .`,
			source:  "a",
			fail:    true,
		},
		{
			name:    "ExtraInputErrors",
			grammar: `Extra = "b" .`,
			source:  "ba",
			tokens:  []string{"b"},
			fail:    true,
		},
		{
			name:    "TokenMatch",
			grammar: `Token = "token" .`,
			source:  `token`,
			tokens:  []string{"token"},
			roots:   []string{"Token"},
		},
		{
			name:    "TokenNoMatch",
			grammar: `Token = "token" .`,
			source:  `toke`,
			fail:    true,
		},
		{
			name:    "RangeMatch",
			grammar: `Range = "a" … "z" .`,
			source:  "x",
			tokens:  []string{"x"},
		},
		{
			name:    "RangeNoMatch",
			grammar: `Range = "a" … "z" .`,
			source:  "A",
			fail:    true,
		},
		{
			name:    "Alternative",
			grammar: `Alternatives = "a" | "b" | "c" .`,
			source:  "a",
			tokens:  []string{"a"},
		},
		{
			name:    "2ndAlternative",
			grammar: `Alternatives = "a" | "b" | "c" .`,
			source:  "b",
			tokens:  []string{"b"},
		},
		{
			name:    "3rdAlternative",
			grammar: `Alternatives = "a" | "b" | "c" .`,
			source:  "c",
			tokens:  []string{"c"},
		},
		{
			name:    "AlternativeDoesNotMatch",
			grammar: `Alternatives = "a" | "b" | "c" .`,
			source:  "d",
			fail:    true,
		},
		{
			name:    "Group",
			grammar: `Group = ("token") .`,
			source:  "token",
			tokens:  []string{"token"},
		},
		{
			name:    "OptionWithInnerMatch",
			grammar: `Option = [ "t" ] .`,
			source:  "t",
			tokens:  []string{"t"},
		},
		{
			name:    "OptionWithNoInnerMatch",
			grammar: `Option = [ "t" ] .`,
			source:  "",
		},
		{
			name: "Identifier",
			grammar: `
			Identifier = alpha { alpha | number } .
			Whitespace = "\n" | "\r" | "\t" | " " .

			alpha = "a"…"z" | "A"…"Z" | "_" .
			number = "0"…"9" .
			`,
			source: `some id withCase andNumb3rs`,
			tokens: []string{"some", " ", "id", " ", "withCase", " ", "andNumb3rs"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defi, err := EBNF(test.grammar)
			if test.failBuild {
				assert.Check(t, err != nil)
				return
			}
			assert.NilError(t, err)
			def := defi.(*ebnfLexerDefinition)
			if test.roots != nil {
				roots := []string{}
				for sym := range def.symbols {
					if sym != "EOF" {
						roots = append(roots, sym)
					}
				}
				assert.DeepEqual(t, test.roots, roots)
			}
			// repr.Println(def, repr.Indent("  "))
			tokens, err := readAllTokens(def.Lex(strings.NewReader(test.source)))
			if test.fail {
				assert.Check(t, err != nil)
			} else {
				assert.NilError(t, err)
			}
			assert.DeepEqual(t, test.tokens, tokens)
		})
	}
}

func readAllTokens(lex Lexer) (out []string, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			if perr, ok := msg.(error); ok {
				err = perr
			} else {
				panic(msg)
			}
		}
	}()
	for {
		token := lex.Next()
		if token.EOF() {
			return
		}
		out = append(out, token.Value)
	}
}

func BenchmarkEBNFLexer(b *testing.B) {
	def, err := EBNF(`
Identifier = alpha { alpha | digit } .
Whitespace = "\n" | "\r" | "\t" | " " .
Number = digit { digit } .

alpha = "a"…"z" | "A"…"Z" | "_" .
digit = "0"…"9" .
`)
	assert.NilError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lex := def.Lex(strings.NewReader("hello world 123 hello world 123"))
		ConsumeAll(lex)
	}
}
