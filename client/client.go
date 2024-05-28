package client

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"net/http"
)

type Client interface {
	Wallet
	SubWallet
}

type client struct {
	domain     string
	apiKey     string
	privateKey *rsa.PrivateKey
	httpClient *http.Client
	RequestID  RequestID
}

type Options struct {
	Domain     string
	HttpClient *http.Client
	RequestID  RequestID
}

func New(apiKey, apiKeySecret string, opts Options) (Client, error) {
	privateKey, err := parseRSAPrivateKey(apiKeySecret)
	if err != nil {
		return nil, err
	}

	if opts.Domain == "" {
		opts.Domain = DefaultDomain
	}
	if opts.HttpClient == nil {
		opts.HttpClient = http.DefaultClient
	}
	if opts.RequestID == nil {
		opts.RequestID = NewRequestID()
	}
	c := &client{
		domain:     opts.Domain,
		apiKey:     apiKey,
		privateKey: privateKey,
		httpClient: opts.HttpClient,
		RequestID:  opts.RequestID,
	}
	return c, nil
}

func parseRSAPrivateKey(privateKeyBase64 string) (*rsa.PrivateKey, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, err
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(decodedPrivateKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
