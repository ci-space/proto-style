package rules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/rule"
)

type EnumInFileEndRule struct {
}

func NewEnumInFileEndRule() *EnumInFileEndRule {
	return &EnumInFileEndRule{}
}

func (r EnumInFileEndRule) ID() string {
	return "PROTOSTYLE_ENUM_IN_FILE_END"
}

func (r EnumInFileEndRule) Purpose() string {
	return "enum must be in file end"
}

func (r EnumInFileEndRule) IsOfficial() bool {
	return false
}

func (r EnumInFileEndRule) Severity() rule.Severity {
	return rule.SeverityError
}

func (r EnumInFileEndRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	itemsCount := 0
	enums := []*parser.Enum{}

	for _, visitee := range proto.ProtoBody {
		switch item := visitee.(type) {
		case *parser.Enum:
			enums = append(enums, item)
		default:
			itemsCount++
		}
	}

	failures := make([]report.Failure, 0)

	for i, enum := range enums {
		goldIndex := itemsCount + i

		if proto.ProtoBody[goldIndex] != enum {
			failures = append(
				failures,
				report.Failuref(enum.Meta.Pos, r.ID(), string(r.Severity()), "Enum %q must be in file end", enum.EnumName),
			)
		}
	}

	return failures, nil
}
