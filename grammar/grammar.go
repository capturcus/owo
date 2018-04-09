package grammar

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
