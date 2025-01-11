package rules

import (
	"fmt"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/rule"
	"github.com/yoheimuta/protolint/linter/visitor"

	"github.com/ci-space/proto-style/internal/utils"
)

type FieldWithOptionRule struct {
	ruleName string
	optName  string
}

type FieldWithOptionVisitor struct {
	*visitor.BaseAddVisitor

	optName string
}

func NewFieldWithOptionRule(
	ruleName string,
	optName string,
) *FieldWithOptionRule {
	return &FieldWithOptionRule{
		ruleName: ruleName,
		optName:  optName,
	}
}

func (r FieldWithOptionRule) ID() string {
	return r.ruleName
}

func (r FieldWithOptionRule) Purpose() string {
	return fmt.Sprintf("Field must have option %q", r.optName)
}

func (r FieldWithOptionRule) IsOfficial() bool {
	return false
}

func (r FieldWithOptionRule) Severity() rule.Severity {
	return rule.SeverityError
}

func (r FieldWithOptionRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &FieldWithOptionVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID(), string(r.Severity())),
		optName:        r.optName,
	}

	return visitor.RunVisitor(v, proto, r.ID())
}

func (v *FieldWithOptionVisitor) VisitField(field *parser.Field) (next bool) {
	if !utils.FieldHasOption(field, v.optName) {
		v.AddFailuref(field.Meta.Pos, "Field %q must have option %q", field.FieldName, v.optName)
	}

	return true
}
