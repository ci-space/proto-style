package utils

import (
	"fmt"

	"github.com/yoheimuta/go-protoparser/v4/parser"
)

func RPCHasOption(rpc *parser.RPC, optName string) bool {
	find := fmt.Sprintf("(%s)", optName)

	for _, option := range rpc.Options {
		if option.OptionName == find {
			return true
		}
	}

	return false
}
