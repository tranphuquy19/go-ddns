package client

import (
	"bytes"
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
	HttpClient *http.Client
}

func InitClient(baseUrl string, tokenChain string, tokenType string) *HttpClient {
	return &HttpClient{
		BaseURL: baseUrl,
		Token: TokenStruct{
			TokenChain: tokenChain,
			TokenType:  tokenType,
		},
		HttpClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

func (c *HttpClient) Get(values []byte, endpoint ...string) (string, error) {
	return worker("GET", c, values, endpoint...)
}

func (c *HttpClient) Post(values []byte, endpoint ...string) (string, error) {
	return worker("POST", c, values, endpoint...)
}

func (c *HttpClient) Del(values []byte, endpoint ...string) (string, error) {
	return worker("DELETE", c, values, endpoint...)
}

func worker(method string, c *HttpClient, values []byte, endpoint ...string) (string, error) {
	url := util.ParseURL(c.BaseURL, endpoint...)

	var (
		req *http.Request
		err error
	)
	if values != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(values))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", c.Token.TokenType, c.Token.TokenChain))

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	bodyStr := string(body)
	// fmt.Println("RESPONSE", bodyStr, "REQUEST", "BODY", string(values), "METHOD", method)

	return bodyStr, err
}
