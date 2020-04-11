package ditto

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		Name, Url, User, Pass string
		Expected              interface{}
		Error                 error
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
			Error: nil,
		},
		{
			Name: "no-url",
			Url:  "",
			User: "ditto",
			Pass: "ditto",
			Expected: Ditto{
				"",
				"ditto",
				"ditto",
			},
			Error: errors.New("invalid url:  "),
		},
		{
			Name: "no-user",
			Url:  "https://ditto.eclipse.org",
			User: "",
			Pass: "ditto",
			Expected: Ditto{
				"https://ditto.eclipse.org",
				"",
				"ditto",
			},
			Error: errors.New("empty user"),
		},
		{
			Name: "no-pass",
			Url:  "https://ditto.eclipse.org",
			User: "ditto",
			Pass: "",
			Expected: Ditto{
				"https://ditto.eclipse.org",
				"ditto",
				"",
			},
			Error: errors.New("empty password"),
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ditto, err := New(test.Url, test.User, test.Pass)
			if err != nil {
				assert.Error(t, err, test.Name)
				return
			}
			assert.NoError(t, err, test.Name)
			assert.Equal(t, test.Expected, ditto, test.Name)
		})
	}
}
