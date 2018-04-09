package grammar

type Program struct {
	Functions []*Function `{ @@ }`
}

type Function struct {
	Identifier string     `@Ident Spaces`
	Args       []*FuncArg `{ @@ } Colon Newline `
	Block      *Block     `@@`
}

type Block struct {
	Statements []*Statement `Indent { @@ Newline } Dedent`
}

type FuncArg struct {
	ArgName string `@Ident Spaces`
	Type    string `@Ident Spaces Comma`
}

type Statement struct {
	Test string `@Ident`
}

/*
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
*/
