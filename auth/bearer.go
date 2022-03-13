package auth

type BearerAuth struct{}

func (auth BearerAuth) parse(authorization string) (token string, ok bool) {
	const prefix = "Bearer "
	if len(authorization) < len(prefix) {
		return
	}

	token = string(authorization[len(prefix):])

	return token, true
}

func (auth BearerAuth) Validate(authorization, token string) (ok bool) {
	if authorization == "" {
		return
	}

	_token, ok := auth.parse(authorization)
	if !ok {
		return
	}

	if token != _token {
		return
	}

	return true
}
