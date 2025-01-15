package utils

import (
	"fmt"
	"strings"
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
		comment.Message = strings.TrimPrefix(comment.Message, "//")
	}

	return comment
}

func (c *Comment) String() string {
	return fmt.Sprintf("%s%s%s", c.Prefix, c.Message, c.Suffix)
}

func (c *Comment) Append(ch string) {
	c.Message += ch
}
