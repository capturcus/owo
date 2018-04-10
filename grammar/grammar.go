package grammar

const LEXER_STRING = `(?m)` +
	`([\t ]+)` +
	`|(?P<Indent>\x11)` +
	`|(?P<Dedent>\x10)` +
	`|(?P<Colon>:)` +
	`|(?P<Newline>\n)` +
	`|(?P<Comma>,)` +
	`|(?P<ExclamMark>!)` +
	`|(?P<Loop>loop)` +
	`|(^[#;].*$)` +
	`|(?P<Ident>[a-zA-Z][a-zA-Z_\d]*)` +
	`|(?P<String>"(?:\\.|[^"])*")` +
	`|(?P<Float>\d+(?:\.\d+)?)` +
	`|(?P<Punct>[][=])`

type Program struct {
	Functions []*Function `{ @@ { Newline } }`
}

type Function struct {
	Identifier string     `@Ident`
	Args       []*FuncArg `[ @@ { Comma @@ } ] Colon Newline`
	Block      *Block     `@@`
}

type Block struct {
	Stmts []*Stmt `Indent { @@ Newline } Dedent`
}

type FuncArg struct {
	ArgName string `@Ident`
	Type    string `@Ident`
}

type Stmt struct {
	Compound *CompoundStmt `@@`
	Simple   *SimpleStmt   `| @@`
	Exclam   bool          `[ @ExclamMark ]`
}

type SimpleStmt struct {
	Pass bool      `@'pass'`
	Flow *FlowStmt `| @@`
	Expr *ExprStmt `| @@`
}

type ExprStmt struct {
	Expr *Expr `@@`
}

type FlowStmt struct {
	Break    bool  `@'break'`
	Continue bool  `| @'continue'`
	Return   *Expr `|'return' @@`
}

type CompoundStmt struct {
	Loop *LoopStmt `@@`
}

type LoopStmt struct {
	Block *Block `Loop Colon Newline @@`
}

type Expr struct {
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
