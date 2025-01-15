package e2e

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func assertLintErrors(t *testing.T, stderr bytes.Buffer, expectedErrors string) {
	expected := strings.Split(strings.TrimSpace(expectedErrors), "\n")

	for _, msg := range expected {
		require.Contains(
			t,
			stderr.String(),
			msg,
			fmt.Sprintf("got %s\n%s", stderr.String(), msg),
		)
	}
}
