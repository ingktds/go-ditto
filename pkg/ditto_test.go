package ditto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		Name, Url, User, Pass string
		Expected              interface{}
	}{
		{
			Name: "normal",
			Url:  "https://ditto.eclipse.org",
			User: "ditto",
			Pass: "ditto",
			Expected: Ditto{
				"https://ditto.eclipse.org",
				"ditto",
				"ditto",
			},
		},
		{
			Name:     "no-url",
			Url:      "",
			User:     "ditto",
			Pass:     "ditto",
			Expected: "ditto.Ditto",
		},
		{
			Name:     "no-user",
			Url:      "https://ditto.eclipse.org",
			User:     "",
			Pass:     "ditto",
			Expected: "ditto.Ditto",
		},
		{
			Name:     "no-pass",
			Url:      "https://ditto.eclipse.org",
			User:     "ditto",
			Pass:     "",
			Expected: "ditto.Ditto",
		},
	}
	for _, test := range tests {
		ditto, err := New(test.Url, test.User, test.Pass)
		if err != nil {
			t.Run(test.Name, func(t *testing.T) {
				assert.IsType(t, test.Expected, ditto)
			})
		} else {
		}
	}
}
