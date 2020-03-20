package ditto

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		Name, Url, User, Pass, Expected string
	}{
		{
			Name:     "normal",
			Url:      "https://ditto.eclipse.org",
			User:     "ditto",
			Pass:     "ditto",
			Expected: "ditto.Ditto",
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
		t.Run(test.Name, func(t *testing.T) {
			if reflect.TypeOf(ditto).String() != test.Expected {
				t.Errorf("Expected: %s", test.Expected)
				t.Errorf("Got: %s", reflect.TypeOf(ditto))
			}
		})
	}
}
