package adminapi

import (
	"fmt"
	"net/url"
)

func Bind(connectURL string) *Client {
	parsed, err := url.Parse(connectURL)
	if err != nil {
		return &Client{
			socketType: "tcp",
			address:    connectURL,
		}
	}

	switch parsed.Scheme {
	case "unix":
		return &Client{
			socketType: "unix",
			address:    parsed.Path,
		}
	case "tcp":
		return &Client{
			socketType: "tcp",
			address:    parsed.Host,
		}
	default:
		panic(fmt.Errorf("unknown scheme %q", parsed.Scheme))
	}
}
