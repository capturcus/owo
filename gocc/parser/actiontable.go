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
			nil,      /* newlines */
			shift(4), /* at */
			shift(5), /* ident */
			nil,      /* colon */
			nil,      /* indent */
			nil,      /* dedent */
			nil,      /* comma */
			nil,      /* empty */
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* newlines */
			shift(4),     /* at */
			shift(5),     /* ident */
			nil,          /* colon */
			nil,          /* indent */
			nil,          /* dedent */
			nil,          /* comma */
			nil,          /* empty */
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(8), /* newlines */
			nil,      /* at */
			nil,      /* ident */
			nil,      /* colon */
			nil,      /* indent */
			nil,      /* dedent */
			nil,      /* comma */
			nil,      /* empty */
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(3), /* newlines, reduce: Declaration */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* newlines */
			nil,      /* at */
			shift(9), /* ident */
			nil,      /* colon */
			nil,      /* indent */
			nil,      /* dedent */
			nil,      /* comma */
			nil,      /* empty */
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			shift(10), /* ident */
			reduce(9), /* colon, reduce: ArgsList */
			nil,       /* indent */
			nil,       /* dedent */
			reduce(9), /* comma, reduce: ArgsList */
			nil,       /* empty */
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(5), /* newlines, reduce: FunctionDeclaration */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(13), /* newlines */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Declarations */
			nil,       /* newlines */
			reduce(1), /* at, reduce: Declarations */
			reduce(1), /* ident, reduce: Declarations */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(14), /* newlines */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			shift(15), /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			nil,       /* ident */
			shift(16), /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			shift(17), /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			nil,       /* ident */
			reduce(7), /* colon, reduce: ArgsList */
			nil,       /* indent */
			nil,       /* dedent */
			reduce(7), /* comma, reduce: ArgsList */
			nil,       /* empty */
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Declarations */
			nil,       /* newlines */
			reduce(2), /* at, reduce: Declarations */
			reduce(2), /* ident, reduce: Declarations */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* newlines */
			nil,      /* at */
			shift(5), /* ident */
			nil,      /* colon */
			nil,      /* indent */
			nil,      /* dedent */
			nil,      /* comma */
			nil,      /* empty */
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* newlines */
			nil,        /* at */
			nil,        /* ident */
			reduce(10), /* colon, reduce: Arg */
			nil,        /* indent */
			nil,        /* dedent */
			reduce(10), /* comma, reduce: Arg */
			nil,        /* empty */
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(19), /* newlines */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			shift(10), /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(4), /* newlines, reduce: FunctionDeclaration */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			shift(21), /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			nil,       /* ident */
			reduce(8), /* colon, reduce: ArgsList */
			nil,       /* indent */
			nil,       /* dedent */
			reduce(8), /* comma, reduce: ArgsList */
			nil,       /* empty */
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			shift(22), /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(13), /* newlines, reduce: Stmt */
			nil,        /* at */
			nil,        /* ident */
			nil,        /* colon */
			nil,        /* indent */
			nil,        /* dedent */
			nil,        /* comma */
			nil,        /* empty */
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* newlines */
			nil,       /* at */
			shift(22), /* ident */
			nil,       /* colon */
			nil,       /* indent */
			shift(25), /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(27), /* newlines */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(6), /* newlines, reduce: BareFunction */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(28), /* newlines */
			nil,       /* at */
			nil,       /* ident */
			nil,       /* colon */
			nil,       /* indent */
			nil,       /* dedent */
			nil,       /* comma */
			nil,       /* empty */
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* newlines */
			nil,        /* at */
			reduce(11), /* ident, reduce: Stmts */
			nil,        /* colon */
			nil,        /* indent */
			reduce(11), /* dedent, reduce: Stmts */
			nil,        /* comma */
			nil,        /* empty */
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* newlines */
			nil,        /* at */
			reduce(12), /* ident, reduce: Stmts */
			nil,        /* colon */
			nil,        /* indent */
			reduce(12), /* dedent, reduce: Stmts */
			nil,        /* comma */
			nil,        /* empty */
		},
	},
}
