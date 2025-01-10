package utils

import (
	"github.com/yoheimuta/go-protoparser/v4/lexer"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/fixer"
)

func ChangeRPCName(rpc *parser.RPC, fix fixer.Fixer, newName string) error {
	const rpcOffset = 3

	namePos := rpc.Meta.Pos
	namePos.Offset += rpcOffset

	return fix.SearchAndReplace(namePos, func(lex *lexer.Lexer) fixer.TextEdit {
		lex.Next()
		return fixer.TextEdit{
			Pos:     lex.Pos.Offset,
			End:     lex.Pos.Offset + len(lex.Text) - 1,
			NewText: []byte(newName),
		}
	})
}

func ChangeRPCRequestName(req *parser.RPCRequest, fix fixer.Fixer, newName string) error {
	const reqOffset = 1

	namePos := req.Meta.Pos
	namePos.Offset += reqOffset

	return fix.SearchAndReplace(namePos, func(lex *lexer.Lexer) fixer.TextEdit {
		lex.Next()
		return fixer.TextEdit{
			Pos:     lex.Pos.Offset,
			End:     lex.Pos.Offset + len(lex.Text) - 1,
			NewText: []byte(newName),
		}
	})
}

func ChangeRPCResponseName(resp *parser.RPCResponse, fix fixer.Fixer, newName string) error {
	const respOffset = 1

	namePos := resp.Meta.Pos
	namePos.Offset += respOffset

	return fix.SearchAndReplace(namePos, func(lex *lexer.Lexer) fixer.TextEdit {
		lex.Next()
		return fixer.TextEdit{
			Pos:     lex.Pos.Offset,
			End:     lex.Pos.Offset + len(lex.Text) - 1,
			NewText: []byte(newName),
		}
	})
}
