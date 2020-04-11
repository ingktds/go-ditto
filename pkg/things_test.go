package ditto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	URL  = "https://ditto.eclipse.org"
	User = "ditto"
	Pass = "ditto"
)

func TestGetThings(t *testing.T) {
	tests := []struct {
		Name     string
		ThingID  string
		Expected Things
	}{
		{
			Name:     "",
			ThingID:  "",
			Expected: Things{},
		},
	}
	d, _ := New(URL, User, Pass)
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got, err := d.GetThings(test.ThingID)
			if err != nil {
				assert.Error(t, err, test.Name)
			}
			assert.NoError(t, err, test.Name)
			assert.Equal(t, test.Expected, got, test.Name)
		})
	}
}
