package proxy

import (
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(targetUrl string) (*httputil.ReverseProxy, error) {
	target, err := url.Parse(targetUrl);
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(target), nil
}