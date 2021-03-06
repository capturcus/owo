// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 2
		case r == 16: // [\u0010,\u0010]
			return 3
		case r == 17: // [\u0011,\u0011]
			return 4
		case r == 32: // [' ',' ']
			return 1
		case r == 44: // [',',',']
			return 5
		case r == 58: // [':',':']
			return 6
		case r == 64: // ['@','@']
			return 7
		case 65 <= r && r <= 90: // ['A','Z']
			return 8
		case r == 95: // ['_','_']
			return 9
		case 97 <= r && r <= 122: // ['a','z']
			return 8
		case r == 243: // [\u00f3,\u00f3]
			return 8
		case r == 261: // [\u0105,\u0105]
			return 8
		case r == 263: // [\u0107,\u0107]
			return 8
		case r == 281: // [\u0119,\u0119]
			return 8
		case r == 322: // [\u0142,\u0142]
			return 8
		case r == 324: // [\u0144,\u0144]
			return 8
		case r == 347: // [\u015b,\u015b]
			return 8
		case r == 378: // [\u017a,\u017a]
			return 8
		case r == 380: // [\u017c,\u017c]
			return 8
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 2
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 95: // ['_','_']
			return 12
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 243: // [\u00f3,\u00f3]
			return 11
		case r == 261: // [\u0105,\u0105]
			return 11
		case r == 263: // [\u0107,\u0107]
			return 11
		case r == 281: // [\u0119,\u0119]
			return 11
		case r == 322: // [\u0142,\u0142]
			return 11
		case r == 324: // [\u0144,\u0144]
			return 11
		case r == 347: // [\u015b,\u015b]
			return 11
		case r == 378: // [\u017a,\u017a]
			return 11
		case r == 380: // [\u017c,\u017c]
			return 11
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 95: // ['_','_']
			return 12
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 243: // [\u00f3,\u00f3]
			return 11
		case r == 261: // [\u0105,\u0105]
			return 11
		case r == 263: // [\u0107,\u0107]
			return 11
		case r == 281: // [\u0119,\u0119]
			return 11
		case r == 322: // [\u0142,\u0142]
			return 11
		case r == 324: // [\u0144,\u0144]
			return 11
		case r == 347: // [\u015b,\u015b]
			return 11
		case r == 378: // [\u017a,\u017a]
			return 11
		case r == 380: // [\u017c,\u017c]
			return 11
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 95: // ['_','_']
			return 12
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 243: // [\u00f3,\u00f3]
			return 11
		case r == 261: // [\u0105,\u0105]
			return 11
		case r == 263: // [\u0107,\u0107]
			return 11
		case r == 281: // [\u0119,\u0119]
			return 11
		case r == 322: // [\u0142,\u0142]
			return 11
		case r == 324: // [\u0144,\u0144]
			return 11
		case r == 347: // [\u015b,\u015b]
			return 11
		case r == 378: // [\u017a,\u017a]
			return 11
		case r == 380: // [\u017c,\u017c]
			return 11
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 95: // ['_','_']
			return 12
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 243: // [\u00f3,\u00f3]
			return 11
		case r == 261: // [\u0105,\u0105]
			return 11
		case r == 263: // [\u0107,\u0107]
			return 11
		case r == 281: // [\u0119,\u0119]
			return 11
		case r == 322: // [\u0142,\u0142]
			return 11
		case r == 324: // [\u0144,\u0144]
			return 11
		case r == 347: // [\u015b,\u015b]
			return 11
		case r == 378: // [\u017a,\u017a]
			return 11
		case r == 380: // [\u017c,\u017c]
			return 11
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 95: // ['_','_']
			return 12
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 243: // [\u00f3,\u00f3]
			return 11
		case r == 261: // [\u0105,\u0105]
			return 11
		case r == 263: // [\u0107,\u0107]
			return 11
		case r == 281: // [\u0119,\u0119]
			return 11
		case r == 322: // [\u0142,\u0142]
			return 11
		case r == 324: // [\u0144,\u0144]
			return 11
		case r == 347: // [\u015b,\u015b]
			return 11
		case r == 378: // [\u017a,\u017a]
			return 11
		case r == 380: // [\u017c,\u017c]
			return 11
		}
		return NoState
	},
}
