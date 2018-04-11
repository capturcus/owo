package ast

import (
	"errors"
	"owo/gocc/token"
	"reflect"
)

type X interface{}

type Declarations struct {
	Datatype *DatatypeDeclaration
	Function *FunctionDeclaration
	Previous *Declarations
}

type DatatypeDeclaration struct {
}

type FunctionDeclaration struct {
	Decorator  string
	Identifier string
	Args       *FuncArg
	Stmts      *Stmt
}

type Block struct {
	Stmts []*Stmt
}

type FuncArg struct {
	ArgName  string
	Type     string
	Previous *FuncArg
}

/*type Stmt struct {
	Compound *CompoundStmt
	Simple   *SimpleStmt
	Exclam   bool
}*/

type Stmt struct {
	Previous *Stmt
	Test     string
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

func TestIdent(d X) (X, error) {
	s := Str(d)
	return &Stmt{Test: s}, nil
}

func Str(x X) string {
	return string(x.(*token.Token).Lit)
}

func AddArg(prev X, curr X) (X, error) {
	curr.(*FuncArg).Previous = prev.(*FuncArg)
	return curr, nil
}

func NewBareFunction(funcName string, args_ X, block_ X) (X, error) {
	args := args_.(*FuncArg)
	block := block_.(*Stmt)
	return &FunctionDeclaration{"", funcName, args, block}, nil
}

func NewAnnotatedFunction(decorator string, fn_ X) (X, error) {
	fn := fn_.(*FunctionDeclaration)
	fn.Decorator = decorator
	return fn, nil
}

func NewDeclaration(dec X) (X, error) {
	switch v := dec.(type) {
	case *DatatypeDeclaration:
		return &Declarations{Datatype: v}, nil
	case *FunctionDeclaration:
		return &Declarations{Function: v}, nil
	default:
		return nil, errors.New("unknown type: " + reflect.ValueOf(dec).String())
	}
}

func AppendDeclaration(prev_ X, curr_ X) (X, error) {
	prev := prev_.(*Declarations)
	curr := curr_.(*Declarations)
	curr.Previous = prev
	return curr, nil
}

func AddStmt(stmts_ X, stmt_ X) (X, error) {
	stmts := stmts_.(*Stmt)
	stmt := stmt_.(*Stmt)
	stmt.Previous = stmts
	return stmt, nil
}
