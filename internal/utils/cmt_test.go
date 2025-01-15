package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnwrapComment(t *testing.T) {
	cases := []struct {
		Raw      string
		Expected *Comment
	}{
		{
			Raw:      "",
			Expected: &Comment{},
		},
		{
			Raw: "// comment",
			Expected: &Comment{
				Prefix:  "//",
				Message: "comment",
			},
		},
		{
			Raw: "/* comment */",
			Expected: &Comment{
				Prefix:  "/*",
				Message: "comment",
				Suffix:  "*/",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Raw, func(t *testing.T) {
			got := UnwrapComment(c.Raw)

			assert.Equal(t, c.Expected, got)
		})
	}
}
