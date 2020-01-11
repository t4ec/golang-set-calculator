package token

type Token struct {
	Type  string
	Value string
}

// Token Types
const (
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"
	WORD     = "WORD"
	EOF      = "EOF"
)
