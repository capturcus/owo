package owoutil

import (
	"fmt"
	"owo/gocc/token"
)

func PrintTok(t *token.Token) {
	fmt.Println("TOKEN", token.TokMap.TokenString(t), t.Lit)
}
