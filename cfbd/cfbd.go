package cfbd

import (
	"net/http"
	"net/url"
	"time"
)

var baseUrl = url.URL{
	Scheme: "https",
	Host:   "api.collegefootballdata.com",
	Path:   "/",
}

type CFBD struct {
	client *http.Client
	token  string
}

func New(token string) *CFBD {
	c := &http.Client{Timeout: time.Minute}

	return &CFBD{
		client: c,
		token:  token,
	}
}
