package utils

import (
	"github.com/yoheimuta/go-protoparser/v4/lexer"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/fixer"
	"regexp"
)

var (
	isListRequestRegex  = regexp.MustCompile("List(.*)Request")
	isListResponseRegex = regexp.MustCompile("List(.*)Response")
)

func ChangeMessageName(rpc *parser.Message, fix fixer.Fixer, newName string) error {
	const msgOffset = 7 // message

	namePos := rpc.Meta.Pos
	namePos.Offset += msgOffset

	return fix.SearchAndReplace(namePos, func(lex *lexer.Lexer) fixer.TextEdit {
		lex.Next()
		return fixer.TextEdit{
			Pos:     lex.Pos.Offset,
			End:     lex.Pos.Offset + len(lex.Text) - 1,
			NewText: []byte(newName),
		}
	})
}

func parseSingleValue(exp *regexp.Regexp, str string) string {
	const needMatchesCount = 2

	matches := exp.FindStringSubmatch(str)
	if len(matches) < needMatchesCount {
		return ""
	}

	return matches[1]
}
