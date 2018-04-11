// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(3), /* ident */
			nil,      /* newline */
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			shift(3),     /* ident */
			nil,          /* newline */
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Declarations */
			reduce(1), /* ident, reduce: Declarations */
			nil,       /* newline */
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Declaration */
			reduce(3), /* ident, reduce: Declaration */
			nil,       /* newline */
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Declarations */
			reduce(2), /* ident, reduce: Declarations */
			nil,       /* newline */
		},
	},
}