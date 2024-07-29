package entities

type Token struct {
	isNative bool
	isToken  bool
}

func NewToken() Token {
	return Token{
		isNative: false,
		isToken:  true,
	}
}
