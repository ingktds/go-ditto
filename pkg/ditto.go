package ditto

import "regexp"

type Ditto struct {
	Url  string
	User string
	Pass string
}

func New(url, user, pass string) (ditto Ditto, err error) {
	if _, err := regexp.MatchString("https?://", url); err != nil {
		return Ditto{}, err
	}
	ditto = Ditto{url, user, pass}
	return ditto, nil
}
