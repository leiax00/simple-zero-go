package server

import (
	"github.com/go-resty/resty/v2"
)

func NewHttpClient() *resty.Client {
	return resty.New()
}
