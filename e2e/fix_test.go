package e2e

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yoheimuta/protolint/lib"
)

func TestFix_Invalid(t *testing.T) {
	cases := []struct {
		RuleName string
	}{
		{
			RuleName: "list_messages_plural_name",
		},
		{
			RuleName: "rpc_without_entity_name",
		},
	}

	for _, c := range cases {
		t.Run(c.RuleName, func(t *testing.T) {
			invalidFile, err := os.ReadFile(fmt.Sprintf("./testdata/%s/invalid.proto", c.RuleName))
			require.NoError(t, err)

			err = os.WriteFile(fmt.Sprintf("./testdata/%s/invalid_fix_run.proto", c.RuleName), invalidFile, 0644)
			require.NoError(t, err)
			defer os.Remove(fmt.Sprintf("./testdata/%s/invalid_fix_run.proto", c.RuleName))

			validProto, err := os.ReadFile(fmt.Sprintf("./testdata/%s/valid.proto", c.RuleName))
			require.NoError(t, err)

			args := []string{
				"-plugin",
				"./protostyle",
				"-fix",
				"-config_path",
				fmt.Sprintf("./testdata/%s/protolint.yaml", c.RuleName),
				fmt.Sprintf("./testdata/%s/invalid_fix_run.proto", c.RuleName),
			}

			var stdout bytes.Buffer
			var stderr bytes.Buffer

			err = lib.Lint(args, &stdout, &stderr)
			require.NoError(t, err)

			got, err := os.ReadFile(fmt.Sprintf("./testdata/%s/invalid_fix_run.proto", c.RuleName))
			require.NoError(t, err)

			assert.Equal(t, string(validProto), string(got))
		})
	}
}
