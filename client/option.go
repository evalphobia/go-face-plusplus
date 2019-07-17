package client

import (
	"time"
)

const (
	baseURL = "https://api-us.faceplusplus.com"
)

// Option contains optional setting of Client.
type Option struct {
	BaseURL   string
	UserAgent string
	Timeout   time.Duration
	Debug     bool
	Retry     bool
}

func (o Option) getBaseURL() string {
	if o.BaseURL != "" {
		return o.BaseURL
	}
	return baseURL
}

func (o Option) getUserAgent() string {
	if o.UserAgent != "" {
		return o.UserAgent
	}
	return "go-face-plusplus"
}

func (o Option) getTimeout() time.Duration {
	if o.Timeout > 0 {
		return o.Timeout
	}
	return 20 * time.Second
}
