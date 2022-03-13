package ip

import (
	"net/http"
)

func Parse(request http.Request) string {
	x_forwarded_for := request.Header.Get("X-Forwarded-For")
	x_real_ip := request.Header.Get("X-Real-IP")
	remote_addr := request.RemoteAddr

	if x_forwarded_for != "" {
		return x_forwarded_for
	}

	if x_real_ip != "" {
		return x_real_ip
	}

	return remote_addr
}
