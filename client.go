package zaico

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://web.zaico.co.jp/api/v1"
)

// Client ZAICO APIクライアント
type Client struct {
	client    *http.Client
	baseURL   *url.URL
	token     string
	UserAgent string

	// サービス
	Inventory *InventoryService
}

// NewClient 新しいZAICO APIクライアントを作成します
func NewClient(token string) *Client {
	return NewClientWithBaseURL(token, defaultBaseURL)
}

// NewClientWithBaseURL 指定されたbaseURLを使用して新しいZAICO APIクライアントを作成します
func NewClientWithBaseURL(token, baseURL string) *Client {
	parsedURL, _ := url.Parse(baseURL)
	c := &Client{
		client:    http.DefaultClient,
		baseURL:   parsedURL,
		token:     token,
		UserAgent: "zaico-go",
	}
	c.Inventory = &InventoryService{client: c}
	return c
}

// NewRequest 新しいAPIリクエストを作成します
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do APIリクエストを実行します
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
