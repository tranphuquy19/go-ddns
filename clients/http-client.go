package client

import (
	"fmt"
	"go-ddns/util"
	"io"
	"net/http"
	"time"
)

type TokenStruct struct {
	TokenChain string
	TokenType  string
}

type HttpClient struct {
	BaseURL    string
	Token      TokenStruct
	HTTPClient *http.Client
}

func InitClient(baseUrl string, tokenChain string, tokenType string) *HttpClient {
	return &HttpClient{
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

func (c *HttpClient) Get(endpoint ...string) (string, error) {
	return worker("GET", c, endpoint...)
}

func (c *HttpClient) Post(endpoint ...string) (string, error) {
	return worker("POST", c, endpoint...)
}

func (c *HttpClient) Del(endpoint ...string) (string, error) {
	return worker("DELETE", c, endpoint...)
}

func worker(method string, c *HttpClient, endpoint ...string) (string, error) {
	url := util.ParseURL(c.BaseURL, endpoint...)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", c.Token.TokenType, c.Token.TokenChain))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	bodyStr := string(body)

	return bodyStr, err
}
