package utils

import (
	"fmt"
	"strings"

	"github.com/yoheimuta/go-protoparser/v4/lexer"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/fixer"
)

type Comment struct {
	Prefix  string
	Message string
	Suffix  string
}

func UnwrapComment(raw string) *Comment {
	comment := &Comment{
		Message: raw,
	}

	if strings.HasPrefix(raw, "/*") {
		comment.Prefix = "/*"
		comment.Suffix = "*/"

		comment.Message = strings.TrimPrefix(comment.Message, "/*")
		comment.Message = strings.TrimSuffix(comment.Message, "*/")
		comment.Message = strings.Trim(comment.Message, " ")
	} else {
		comment.Suffix = "//"
		comment.Message = strings.TrimPrefix(comment.Message, "//")
	}

	return comment
}

func (c *Comment) String() string {
	return fmt.Sprintf("%s%s%s", c.Suffix, c.Message, c.Prefix)
}

func (c *Comment) Append(ch string) {
	c.Message += ch
}

func ChangeComment(cmt *parser.Comment, fix fixer.Fixer, newComment string) error {
	pos := cmt.Meta.Pos

	return fix.SearchAndReplace(pos, func(lex *lexer.Lexer) fixer.TextEdit {
		return fixer.TextEdit{
			Pos:     lex.Pos.Offset,
			End:     lex.Pos.Offset + len(cmt.Raw) - 1,
			NewText: []byte(newComment),
		}
	})
}
