package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yoheimuta/protolint/lib"
)

func TestLint_Invalid(t *testing.T) {
	cases := []struct {
		RuleName string
	}{
		{
			RuleName: "field_with_behavior",
		},
		{
			RuleName: "list_messages_resource_name_pluralized",
		},
		{
			RuleName: "rpc_without_resource_name",
		},
	}

	for _, c := range cases {
		t.Run(c.RuleName, func(t *testing.T) {
			invalidMessage, err := os.ReadFile(fmt.Sprintf("./testdata/%s/invalid_message.txt", c.RuleName))
			require.NoError(t, err)

			args := []string{
				"-plugin",
				"./protostyle",
				"-config_path",
				fmt.Sprintf("./testdata/%s/protolint.yaml", c.RuleName),
				fmt.Sprintf("./testdata/%s/invalid.proto", c.RuleName),
			}

			var stdout bytes.Buffer
			var stderr bytes.Buffer

			err = lib.Lint(args, &stdout, &stderr)
			require.Error(t, err, "lint error")

			assertLintErrors(t, stderr, string(invalidMessage))
		})
	}
}
