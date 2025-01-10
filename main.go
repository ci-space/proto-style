package main

import (
	"github.com/yoheimuta/protolint/linter/rule"
	"github.com/yoheimuta/protolint/plugin"

	"github.com/ci-space/proto-style/internal/rules"
)

func main() {
	plugin.RegisterCustomRules(
		plugin.RuleGen(func(_ bool, fixMode bool) rule.Rule {
			return rules.NewRPCWithoutServiceNameRule(fixMode)
		}),
		rules.NewRPCWithHTTP(),
		plugin.RuleGen(func(_ bool, fixMode bool) rule.Rule {
			return rules.NewListMessagesPluralNameRule(fixMode)
		}),
	)
}
