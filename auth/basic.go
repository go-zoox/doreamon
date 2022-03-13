package auth

import (
	"encoding/base64"
	"strings"
)

type BasicAuth struct{}

func (auth BasicAuth) Validate(authorization, username, password string) (ok bool) {
	if authorization == "" {
		return
	}

	_username, _password, ok := auth.parse(authorization)
	if !ok {
		return
	}

	if username != _username || password != _password {
		return
	}

	return true
}

func (auth BasicAuth) parse(authorization string) (username, password string, ok bool) {
	const prefix = "Basic "
	if len(authorization) < len(prefix) || !strings.EqualFold(authorization[:len(prefix)], prefix) {
		return
	}

	c, err := base64.StdEncoding.DecodeString(authorization[len(prefix):])
	if err != nil {
		return
	}

	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}

	return cs[:s], cs[s+1:], true
}
