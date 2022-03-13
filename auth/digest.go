package auth

import (
	"fmt"
	"strings"
)

type DigestAuth struct {
}

type DigestAuthChallenge struct {
	username  string
	realm     string
	nonce     string
	uri       string
	algorithm string
	response  string
	opaque    string
	qop       string
	nc        string
	cnonce    string
}

type DigestAuthChallengeMap map[string]string

func (d DigestAuthChallengeMap) Get(key string) string {
	return d[key]
}

func (auth DigestAuth) parse(authorization string) (challenge DigestAuthChallenge, ok bool) {
	_info := make(DigestAuthChallengeMap)

	const prefix = "Digest "
	if len(authorization) < len(prefix) {
		return
	}

	data := authorization[len(prefix):]
	list := strings.Split(data, ",")
	for _, one := range list {
		kv := strings.Split(one, "=")
		key := strings.Trim(kv[0], " ")
		value := strings.Trim(kv[1], " ")
		if f := strings.IndexByte(kv[1], '"'); f != -1 {
			value = strings.Split(kv[1], "\"")[1]
		}

		// fmt.Printf("found: %s = %s(%s)\n", key, value, reflect.TypeOf(value))
		_info[key] = value
	}

	challenge.username = _info.Get("username")
	challenge.realm = _info.Get("realm")
	challenge.nonce = _info.Get("nonce")
	challenge.uri = _info.Get("uri")
	challenge.algorithm = _info.Get("algorithm")
	challenge.qop = _info.Get("qop")
	challenge.response = _info.Get("response")
	challenge.opaque = _info.Get("opaque")
	challenge.nc = _info.Get("nc")
	challenge.cnonce = _info.Get("cnonce")

	return challenge, true
}

func (auth DigestAuth) Validate(authorization string, qop, username, password, method, path string) (ok bool) {
	if authorization == "" {
		return
	}

	challenge, ok := auth.parse(authorization)
	if !ok {
		return
	}

	// fmt.Printf("%s:%s:%s", challenge.username, challenge.realm, password)
	// fmt.Printf("%s:%s\n", method, path)

	ha1 := md5.Md5(fmt.Sprintf("%s:%s:%s", challenge.username, challenge.realm, password))
	ha2 := md5.Md5(fmt.Sprintf("%s:%s", method, path))
	response := md5.Md5(fmt.Sprintf(
		"%s:%s:%s:%s:%s:%s",
		ha1,
		challenge.nonce,
		challenge.nc,
		challenge.cnonce,
		qop,
		ha2,
	))

	// fmt.Printf(
	// 	"response text: %s:%s:%s:%s:%s:%s\n",
	// 	ha1,
	// 	challenge.nonce,
	// 	challenge.nc,
	// 	challenge.cnonce,
	// 	qop,
	// 	ha2,
	// )
	// fmt.Printf("response: client(%s) server(%s)\n", challenge.response, response)

	return response == challenge.response
}
