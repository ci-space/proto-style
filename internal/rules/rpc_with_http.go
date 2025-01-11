package rules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/rule"
	"github.com/yoheimuta/protolint/linter/visitor"

	"github.com/ci-space/protostyle/internal/utils"
)

const optionHTTPName = "google.api.http"

type RPCWithHTTPRule struct {
}

type rpcWithHTTPVisitor struct {
	*visitor.BaseAddVisitor
}

func NewRPCWithHTTP() *RPCWithHTTPRule {
	return &RPCWithHTTPRule{}
}

func (r RPCWithHTTPRule) ID() string {
	return "PROTOSTYLE_RPC_WITH_HTTP"
}

func (r RPCWithHTTPRule) Purpose() string {
	return "method must have http option"
}

func (r RPCWithHTTPRule) IsOfficial() bool {
	return false
}

func (r RPCWithHTTPRule) Severity() rule.Severity {
	return rule.SeverityError
}

func (r RPCWithHTTPRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &rpcWithHTTPVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID(), string(r.Severity())),
	}

	return visitor.RunVisitor(v, proto, r.ID())
}

func (v *rpcWithHTTPVisitor) VisitService(srv *parser.Service) (next bool) {
	for _, srvBody := range srv.ServiceBody {
		srvRPC, isSrvRPC := srvBody.(*parser.RPC)
		if !isSrvRPC {
			continue
		}

		if !utils.RPCHasOption(srvRPC, optionHTTPName) {
			v.AddFailuref(
				srvRPC.Meta.Pos,
				"%s.%s: Procedure must have http option %q",
				srv.ServiceName,
				srvRPC.RPCName,
				optionHTTPName,
			)
		}
	}

	return true
}
