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
	} else if strings.HasPrefix(raw, "//") {
		comment.Prefix = "//"
		comment.Message = strings.TrimPrefix(comment.Message, "//")
	}

	comment.Message = strings.Trim(comment.Message, " ")

	return comment
}

func (c *Comment) String() string {
	if c.Prefix != "" {
		if c.Suffix != "" {
			return fmt.Sprintf("%s %s %s", c.Prefix, c.Message, c.Suffix)
		}

		return fmt.Sprintf("%s %s", c.Prefix, c.Message)
	}

	return c.Message
}

func (c *Comment) Append(ch string) {
	c.Message += ch
}

func (c *Comment) HasSuffix() bool {
	return c.Suffix != ""
}

func ChangeComment(cmt *parser.Comment, fix fixer.Fixer, newComment *Comment) error {
	start := cmt.Meta.Pos

	newCmt := newComment.String()
	if cmt.Meta.Pos.Column <= 1 && !newComment.HasSuffix() {
		start.Offset--
		newCmt = "/ " + newComment.Message + "\n"
	}

	return fix.SearchAndReplace(start, func(lex *lexer.Lexer) fixer.TextEdit {
		return fixer.TextEdit{
			Pos:     lex.Pos.Offset,
			End:     lex.Pos.Offset + len(cmt.Raw) - 1,
			NewText: []byte(newCmt),
		}
	})
}
