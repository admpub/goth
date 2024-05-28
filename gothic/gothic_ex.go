package gothic

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	redirectURIQueryString = `redirect_uri=%2F`
	HeaderXForwardedHost   = `X-Forwarded-Host`
	HeaderXForwardedProto  = `X-Forwarded-Proto`
	SchemeHTTPS            = `https`
	SchemeHTTP             = `http`
)

func CurrentRootURL(req *http.Request) string {
	var host string
	var scheme string
	if xfh := req.Header.Get(HeaderXForwardedHost); len(xfh) == 0 {
		host = req.Host
	} else {
		host = xfh
	}
	if xfp := req.Header.Get(HeaderXForwardedProto); len(xfp) == 0 {
		if req.TLS != nil {
			scheme = SchemeHTTPS
		} else {
			scheme = SchemeHTTP
		}
	} else {
		scheme = xfp
	}
	return scheme + `://` + host + `/`
}

func FixedRedirectURIQueryString(req *http.Request, authURL string) string {
	pos := strings.Index(authURL, redirectURIQueryString)
	if pos > 0 {
		authURL = authURL[0:pos] + `redirect_uri=` + url.QueryEscape(CurrentRootURL(req)) + authURL[pos+len(redirectURIQueryString):]
	}
	return authURL
}
