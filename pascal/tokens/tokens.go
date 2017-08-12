package tokens

import (
	"strings"
	"unicode"
	"github.com/sunshower-io/updraft/common/io"
	"github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/tokens"
)


var (
	RESERVED_WORDS map[string]core.TokenType = make(map[string]core.TokenType)

	AND       = core.ReservedWord("AND", RESERVED_WORDS)
	ARRAY     = core.ReservedWord("ARRAY", RESERVED_WORDS)
	BEGIN     = core.ReservedWord("BEGIN", RESERVED_WORDS)
	CASE      = core.ReservedWord("CASE", RESERVED_WORDS)
	CONST     = core.ReservedWord("CONST", RESERVED_WORDS)
	DIV       = core.ReservedWord("DIV", RESERVED_WORDS)
	DO        = core.ReservedWord("DO", RESERVED_WORDS)
	DOWNTO    = core.ReservedWord("DOWNTO", RESERVED_WORDS)
	ELSE      = core.ReservedWord("ELSE", RESERVED_WORDS)
	END       = core.ReservedWord("END", RESERVED_WORDS)
	FILE      = core.ReservedWord("FILE", RESERVED_WORDS)
	FOR       = core.ReservedWord("FOR", RESERVED_WORDS)
	FUNCTION  = core.ReservedWord("FUNCTION", RESERVED_WORDS)
	GOTO      = core.ReservedWord("GOTO", RESERVED_WORDS)
	IF        = core.ReservedWord("IF", RESERVED_WORDS)
	IN        = core.ReservedWord("IN", RESERVED_WORDS)
	LABEL     = core.ReservedWord("LABEL", RESERVED_WORDS)
	MOD       = core.ReservedWord("MOD", RESERVED_WORDS)
	NIL       = core.ReservedWord("NIL", RESERVED_WORDS)
	NOT       = core.ReservedWord("NOT", RESERVED_WORDS)
	OF        = core.ReservedWord("OF", RESERVED_WORDS)
	OR        = core.ReservedWord("OR", RESERVED_WORDS)
	PACKED    = core.ReservedWord("PACKED", RESERVED_WORDS)
	PROCEDURE = core.ReservedWord("PROCEDURE", RESERVED_WORDS)
	PROGRAM   = core.ReservedWord("PROGRAM", RESERVED_WORDS)
	RECORD    = core.ReservedWord("RECORD", RESERVED_WORDS)
	REPEAT    = core.ReservedWord("REPEAT", RESERVED_WORDS)
	SET       = core.ReservedWord("SET", RESERVED_WORDS)
	THEN      = core.ReservedWord("THEN", RESERVED_WORDS)
	TO        = core.ReservedWord("TO", RESERVED_WORDS)
	TYPE      = core.ReservedWord("TYPE", RESERVED_WORDS)
	UNTIL     = core.ReservedWord("UNTIL", RESERVED_WORDS)
	VAR       = core.ReservedWord("VAR", RESERVED_WORDS)
	WHILE     = core.ReservedWord("WHILE", RESERVED_WORDS)
	WITH      = core.ReservedWord("WITH", RESERVED_WORDS)

	IDENTIFIER = core.CreateTT("IDENTIFIER")
	INTEGER    = core.CreateTT("INTEGER")
	REAL       = core.CreateTT("REAL")
	STRING     = core.CreateTT("STRING")
	EOF        = core.CreateTT("EOF")

	// Special Symbols

	SPECIAL_TOKENS map[string]core.TokenType = make(map[string]core.TokenType)

	PLUS         = core.SymbolTokenType("PLUS", "+", SPECIAL_TOKENS)
	MINUS        = core.SymbolTokenType("MINUS", "-", SPECIAL_TOKENS)
	STAR         = core.SymbolTokenType("STAR", "*", SPECIAL_TOKENS)
	SLASH        = core.SymbolTokenType("SLASH", "/", SPECIAL_TOKENS)
	COLON_EQUALS = core.SymbolTokenType("ASSIGNMENT", ":=", SPECIAL_TOKENS)
	DOT          = core.SymbolTokenType("DOT", ".", SPECIAL_TOKENS)
	COMMMA       = core.SymbolTokenType("COMMA", ",", SPECIAL_TOKENS)
	SEMICOLON    = core.SymbolTokenType("SEMICOLON", ";", SPECIAL_TOKENS)
	COLON        = core.SymbolTokenType("COLON", ":", SPECIAL_TOKENS)
	QUOTE        = core.SymbolTokenType("QUOTE", "'", SPECIAL_TOKENS)
	EQUALS       = core.SymbolTokenType("EQUALS", "=", SPECIAL_TOKENS)
	NOT_EQUALS   = core.SymbolTokenType("NOT_EQUALS", "<>", SPECIAL_TOKENS)
	LT           = core.SymbolTokenType("LT", "<", SPECIAL_TOKENS)
	LTE          = core.SymbolTokenType("LTE", "<=", SPECIAL_TOKENS)
	GT           = core.SymbolTokenType("GT", ">", SPECIAL_TOKENS)
	GTE          = core.SymbolTokenType("GTE", ">=", SPECIAL_TOKENS)
	LPAREN       = core.SymbolTokenType("LPAREN", "(", SPECIAL_TOKENS)
	RPAREN       = core.SymbolTokenType("RPAREN", ")", SPECIAL_TOKENS)
	LBRACKET     = core.SymbolTokenType("LBRACKET", "[", SPECIAL_TOKENS)
	RBRACKET     = core.SymbolTokenType("RBRACKET", "]", SPECIAL_TOKENS)
	LBRACE       = core.SymbolTokenType("LBRACE", "{", SPECIAL_TOKENS)
	RBRACE       = core.SymbolTokenType("RBRACE", "}", SPECIAL_TOKENS)
	CARET        = core.SymbolTokenType("CARET", "^", SPECIAL_TOKENS)
	DOT_DOT      = core.SymbolTokenType("DOT_DOT", "..", SPECIAL_TOKENS)
)

func IsSymbol(ch rune) bool {
	str := string(ch)
	_, exists := SPECIAL_TOKENS[str]
	return exists
}

func ReservedWord(w string) (core.TokenType, bool) {
	ttype, exists := RESERVED_WORDS[w]
	return ttype, exists
}

func LookupReservedWord(uppercaseName string) (core.TokenType, bool) {
	a, l := RESERVED_WORDS[uppercaseName]
	return a, l
}

type PascalToken struct {
	*core.BaseToken
}

type ErrorToken struct {
	*core.BaseToken
	Code tokens.ErrorCode
}

func (t *PascalToken) Extract() error {
	buffer := make([]rune, 0)
	var (
		ch rune
		er error
	)
	source := t.Source

	for ch, er = source.CurrentCharacter(); er == nil; {
		if unicode.IsSpace(ch) {
			ch, er = source.NextCharacter()
			break
		}
		if unicode.IsPunct(ch) {
			break
		} else {
			buffer = append(buffer, ch)
			ch, er = source.NextCharacter()
		}
	}
	value := string(buffer)
	reservedWord, exists := LookupReservedWord(strings.ToUpper(value))
	if exists {
		t.Type = reservedWord
	} else {
		t.Type = IDENTIFIER
	}
	t.Text = value
	return er
}

func NewPascalToken(t core.TokenType, s io.Source) (core.Token, error) {
	pt := &PascalToken{
		BaseToken: core.NewToken(s, t).(*core.BaseToken),
	}
	pt.Extract()
	return pt, nil
}


func NewError(
		source io.Source,
		code tokens.ErrorCode,
		value string,
) core.Token {
	
	
	
	return &ErrorToken {
		Code		: code,
		BaseToken	:  core.CreateToken(
			source,
			core.ERROR_TOKEN,
			string(code),
			value,
		).(*core.BaseToken),
	}
}