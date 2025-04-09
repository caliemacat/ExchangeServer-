package client

import (
	"crypto/tls"
	"github.com/imroc/req/v3"
)

func Global() *req.Client {
	var Client = req.C()
	Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	Client.TLSClientConfig.InsecureSkipVerify = true
	return Client
}
