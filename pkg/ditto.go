package ditto

import (
	"fmt"
	"regexp"
)

type Ditto struct {
	Url  string
	User string
	Pass string
}

func New(url, user, pass string) (ditto Ditto, err error) {
	ret, err := regexp.MatchString("https?://", url)
	if err != nil {
		return ditto, fmt.Errorf("invalid url: %s", url)
	}
	if !ret {
		return ditto, fmt.Errorf("invalid url: %s", url)
	}
	if user == "" {
		return ditto, fmt.Errorf("empty user")
	}
	if pass == "" {
		return ditto, fmt.Errorf("empty password")
	}
	ditto = Ditto{url, user, pass}
	return ditto, err
}
