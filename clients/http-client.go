package client

import (
	"net/http"
	"time"
)

type TokenStruct struct {
	TokenChain string
	TokenType  string
}

type Client struct {
	BaseURL    string
	Token      TokenStruct
	HTTPClient *http.Client
}

func InitClient(baseUrl string, tokenChain string, tokenType string) *Client {
	return &Client{
		BaseURL: baseUrl,
		Token: TokenStruct{
			TokenChain: tokenChain,
			TokenType:  tokenType,
		},
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}
