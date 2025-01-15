package rules

import (
	"fmt"
	"strings"

	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/rule"
	"github.com/yoheimuta/protolint/linter/visitor"

	"github.com/ci-space/protostyle/internal/utils"
)

type RCPWithoutResourceNameRule struct {
	fixMode bool
}

type rpcWithoutResourceNameVisitor struct {
	*visitor.BaseFixableVisitor

	fixMode bool
}

func NewRPCWithoutServiceNameRule(fixMode bool) *RCPWithoutResourceNameRule {
	return &RCPWithoutResourceNameRule{fixMode: fixMode}
}

func (r RCPWithoutResourceNameRule) ID() string {
	return "PROTOSTYLE_RPC_WITHOUT_RESOURCE_NAME"
}

func (r RCPWithoutResourceNameRule) Purpose() string {
	return "method name must not contain resource name"
}

func (r RCPWithoutResourceNameRule) IsOfficial() bool {
	return false
}

func (r RCPWithoutResourceNameRule) Severity() rule.Severity {
	return rule.SeverityError
}

func (r RCPWithoutResourceNameRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	baseVisitor, err := visitor.NewBaseFixableVisitor(r.ID(), r.fixMode, proto, string(r.Severity()))
	if err != nil {
		return nil, fmt.Errorf("failed to create base fixable visitor: %w", err)
	}

	v := &rpcWithoutResourceNameVisitor{
		BaseFixableVisitor: baseVisitor,
		fixMode:            r.fixMode,
	}

	return visitor.RunVisitor(v, proto, r.ID())
}

func (v *rpcWithoutResourceNameVisitor) VisitService(srv *parser.Service) (next bool) {
	for _, srvBody := range srv.ServiceBody {
		srvRPC, isSrvRPC := srvBody.(*parser.RPC)
		if !isSrvRPC {
			continue
		}

		entName := utils.ParseResourceNameFromServiceName(srv.ServiceName)
		if strings.HasSuffix(srvRPC.RPCName, entName.Singular) || strings.HasSuffix(srvRPC.RPCName, entName.Plural) {
			expectedName := v.createExpectedName(srvRPC.RPCName, entName)

			if v.fixMode {
				err := utils.ChangeRPCName(srvRPC, v.Fixer, expectedName)
				if err != nil {
					panic(err)
				}
			} else {
				v.AddFailuref(
					srvRPC.Meta.Pos,
					"%s.%s: The name of the RPC %q must not contain the name of the service (%s|%s). Expected name: %s",
					srv.ServiceName,
					srvRPC.RPCName,
					srvRPC.RPCName,
					entName.Singular,
					entName.Plural,
					expectedName,
				)
			}
		}
	}

	return true
}

func (v *rpcWithoutResourceNameVisitor) createExpectedName(rpcName string, entName utils.ResourceName) string {
	expected := rpcName
	expected = strings.ReplaceAll(expected, entName.Plural, "")
	expected = strings.ReplaceAll(expected, entName.Singular, "")
	return expected
}
