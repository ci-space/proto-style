package rules

import (
	"fmt"
	"github.com/ci-space/protostyle/internal/utils"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/rule"
	"github.com/yoheimuta/protolint/linter/visitor"
	"strings"
)

type CommentEndsDotRule struct {
	fixMode bool
}

type commentEndsDotVisitor struct {
	*visitor.BaseFixableVisitor

	fixMode bool
}

func NewCommentEndsDotRule(fixMode bool) *CommentEndsDotRule {
	return &CommentEndsDotRule{fixMode: fixMode}
}

func (r CommentEndsDotRule) ID() string {
	return "PROTOSTYLE_COMMENT_ENDS_DOT"
}

func (r CommentEndsDotRule) Purpose() string {
	return "comment must ends dot"
}

func (r CommentEndsDotRule) IsOfficial() bool {
	return false
}

func (r CommentEndsDotRule) Severity() rule.Severity {
	return rule.SeverityError
}

func (r CommentEndsDotRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	baseVisitor, err := visitor.NewBaseFixableVisitor(r.ID(), r.fixMode, proto, string(r.Severity()))
	if err != nil {
		return nil, fmt.Errorf("failed to create base fixable visitor: %w", err)
	}

	v := &commentEndsDotVisitor{
		BaseFixableVisitor: baseVisitor,
		fixMode:            r.fixMode,
	}

	return visitor.RunVisitor(v, proto, r.ID())
}

func (v *commentEndsDotVisitor) VisitComment(c *parser.Comment) {
	comment := utils.UnwrapComment(c.Raw)

	if strings.HasSuffix(comment.Message, ".") {
		return
	}

	if !v.fixMode {
		v.AddFailuref(c.Meta.Pos, "Comment must ends dot")
		return
	}

	comment.Append(".")

	if err := utils.ChangeComment(c, v.Fixer, comment); err != nil {
		panic(err)
	}
}
