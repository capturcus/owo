package lexer

import (
	"strings"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
)

func TestMap(t *testing.T) {
	def := Must(Regexp(`(?P<Whitespace>\s+)|(?P<Ident>\w+)`))

	// Remove whitespace and upper case all other tokens.
	mapper := Map(def, func(t *Token) *Token {
		if t.Type == def.Symbols()["Whitespace"] {
			return nil
		}
		t.Value = strings.ToUpper(t.Value)
		return t
	})

	mappingLexer := mapper.Lex(strings.NewReader("hello world"))
	actual, err := ConsumeAll(mappingLexer)
	assert.NilError(t, err)

	expected := []Token{
		{Type: -3, Value: "HELLO", Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1}},
		{Type: -3, Value: "WORLD", Pos: Position{Filename: "", Offset: 6, Line: 1, Column: 7}},
		{Type: -1, Value: "<<EOF>>", Pos: Position{Filename: "", Offset: 11, Line: 1, Column: 12}},
	}

	assert.DeepEqual(t, expected, actual)
}

func TestUnquote(t *testing.T) {
	def := Unquote(Must(Regexp(`(\s+)|(?P<Ident>\w+)|(?P<String>"[^"]+")`)))
	lexer := def.Lex(strings.NewReader(`hello "world"`))
	actual, err := ConsumeAll(lexer)
	assert.NilError(t, err)
	expected := []Token{
		{Type: -3, Value: "hello", Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1}},
		{Type: -4, Value: "world", Pos: Position{Filename: "", Offset: 6, Line: 1, Column: 7}},
		{Type: -1, Value: "<<EOF>>", Pos: Position{Filename: "", Offset: 13, Line: 1, Column: 14}},
	}
	assert.DeepEqual(t, expected, actual)
}

func TestUnquoteSingleQuote(t *testing.T) {
	def := Unquote(Must(Regexp(`(\s+)|(?P<Ident>\w+)|(?P<String>'(\\.|[^'])*'|"[^"]*")`)))
	lexer := def.Lex(strings.NewReader(`hello 'world\''`))
	actual, err := ConsumeAll(lexer)
	assert.NilError(t, err)
	expected := []Token{
		{Type: -3, Value: "hello", Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1}},
		{Type: -4, Value: "world'", Pos: Position{Filename: "", Offset: 6, Line: 1, Column: 7}},
		{Type: -1, Value: "<<EOF>>", Pos: Position{Filename: "", Offset: 15, Line: 1, Column: 16}},
	}
	assert.DeepEqual(t, expected, actual)
}

func TestElide(t *testing.T) {
	def := Elide(Must(Regexp(`(?P<Whitespace>\s+)|(?P<Ident>\w+)`)), "Whitespace")
	lex := def.Lex(strings.NewReader("hello world"))
	actual, err := ConsumeAll(lex)
	assert.NilError(t, err)
	expected := []Token{
		{Type: -3, Value: "hello", Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1}},
		{Type: -3, Value: "world", Pos: Position{Filename: "", Offset: 6, Line: 1, Column: 7}},
		{Type: -1, Value: "<<EOF>>", Pos: Position{Filename: "", Offset: 11, Line: 1, Column: 12}},
	}
	assert.DeepEqual(t, expected, actual)
}
