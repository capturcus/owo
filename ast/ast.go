package ast

type X interface{}

type Declarations struct {
	/*Datatype        *DatatypeDeclaration
	Function        *FunctionDeclaration*/
	Text string
	Next *Declarations
}

type DatatypeDeclaration struct {
}

type FunctionDeclaration struct {
	Annotated *AnnotatedFunction
	Bare      *BareFunction
}

type AnnotatedFunction struct {
	Decorator  string
	Identifier string
	Args       []*FuncArg
	Block      *Block
}

type BareFunction struct {
	Identifier string
	Args       []*FuncArg
	Block      *Block
}

type Block struct {
	Stmts []*Stmt
}

type FuncArg struct {
	ArgName string
	Type    string
}

type Stmt struct {
	Compound *CompoundStmt
	Simple   *SimpleStmt
	Exclam   bool
}

type SimpleStmt struct {
	Pass   bool
	Assign *AssignStmt
	Flow   *FlowStmt
	Expr   *ExprStmt
}

type AssignStmt struct {
	Locations []*Location
	Exprs     []*Expr
}

type Location struct {
	Var string
}

type ExprStmt struct {
	Expr *Expr
}

type FlowStmt struct {
	Break    bool
	Continue bool
	Return   *Expr
}

type CompoundStmt struct {
	Loop *LoopStmt
}

type LoopStmt struct {
	Block *Block
}

type Expr struct {
	Variable string
}
