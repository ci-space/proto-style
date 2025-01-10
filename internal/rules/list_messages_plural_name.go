package rules

import (
	"fmt"

	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/rule"
	"github.com/yoheimuta/protolint/linter/visitor"

	"github.com/ci-space/proto-style/internal/utils"
)

type ListMessagesPluralNameRule struct {
	fixMode bool
}

type listMessagesPluralNameVisitor struct {
	*visitor.BaseFixableVisitor

	fixMode bool
}

func NewListMessagesPluralNameRule(fixMode bool) *ListMessagesPluralNameRule {
	return &ListMessagesPluralNameRule{fixMode: fixMode}
}

func (r ListMessagesPluralNameRule) ID() string {
	return "PROTOSTYLE_LIST_MESSAGES_PLURAL_NAME_RULE"
}

func (r ListMessagesPluralNameRule) Purpose() string {
	return "list request/response must have plural entity name"
}

func (r ListMessagesPluralNameRule) IsOfficial() bool {
	return false
}

func (r ListMessagesPluralNameRule) Severity() rule.Severity {
	return rule.SeverityError
}

func (r ListMessagesPluralNameRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	baseVisitor, err := visitor.NewBaseFixableVisitor(r.ID(), r.fixMode, proto, string(r.Severity()))
	if err != nil {
		return nil, fmt.Errorf("failed to create base fixable visitor: %w", err)
	}

	v := &listMessagesPluralNameVisitor{
		BaseFixableVisitor: baseVisitor,
		fixMode:            r.fixMode,
	}

	return visitor.RunVisitor(v, proto, r.ID())
}

func (v *listMessagesPluralNameVisitor) VisitRPC(rpc *parser.RPC) (next bool) {
	next = v.visitRPCRequest(rpc.RPCRequest)
	if !next {
		return false
	}

	return v.visitRPCResponse(rpc.RPCResponse)
}

func (v *listMessagesPluralNameVisitor) visitRPCRequest(req *parser.RPCRequest) (next bool) {
	entName := utils.ParseEntityNameFromListRequestName(req.MessageType)
	if entName.Singular == "" || entName.IsPlural() {
		return true
	}

	expectedName := fmt.Sprintf("List%sRequest", entName.Plural)

	if v.fixMode {
		err := utils.ChangeRPCRequestName(req, v.Fixer, expectedName)
		if err != nil {
			panic(err)
		}
	} else {
		v.AddFailuref(
			req.Meta.Pos,
			"%s: must be have plural entity name. Expected name: %q",
			req.MessageType,
			expectedName,
		)
	}

	return true
}

func (v *listMessagesPluralNameVisitor) visitRPCResponse(resp *parser.RPCResponse) (next bool) {
	entName := utils.ParseEntityNameFromListResponseName(resp.MessageType)
	if entName.Singular == "" || entName.IsPlural() {
		return true
	}

	expectedName := fmt.Sprintf("List%sResponse", entName.Plural)

	if v.fixMode {
		err := utils.ChangeRPCResponseName(resp, v.Fixer, expectedName)
		if err != nil {
			panic(err)
		}
	} else {
		v.AddFailuref(
			resp.Meta.Pos,
			"%s: must be have plural entity name. Expected name: %q",
			resp.MessageType,
			expectedName,
		)
	}

	return true
}

func (v *listMessagesPluralNameVisitor) VisitMessage(msg *parser.Message) (next bool) {
	isReq := true

	entName := utils.ParseEntityNameFromListRequestName(msg.MessageName)
	if entName.Singular == "" {
		isReq = false
		entName = utils.ParseEntityNameFromListResponseName(msg.MessageName)
	}
	if entName.Singular == "" {
		return true
	}

	if !entName.IsPlural() {
		expectedName := "List" + entName.Plural
		if isReq {
			expectedName += "Request"
		} else {
			expectedName += "Response"
		}

		if v.fixMode {
			err := utils.ChangeMessageName(msg, v.Fixer, expectedName)
			if err != nil {
				panic(err)
			}
		} else {
			v.AddFailuref(
				msg.Meta.Pos,
				"%s: must be have plural entity name. Expected name: %q",
				msg.MessageName,
				expectedName,
			)
		}
	}

	return true
}
