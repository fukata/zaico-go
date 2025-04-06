package main 

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// setup テスト用のモックサーバーとクライアントを設定します
func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClientWithBaseURL("test-token", server.URL)

	return client, mux, server.Close
}

// testMethod リクエストのメソッドを検証します
func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

// testHeader リクエストのヘッダーを検証します
func testHeader(t *testing.T, r *http.Request, header string, want string) {
	t.Helper()
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

// testBody リクエストのボディを検証します
func testBody(t *testing.T, r *http.Request, want interface{}) {
	t.Helper()
	var v interface{}
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		t.Fatalf("decode json: %v", err)
	}
	if !reflect.DeepEqual(v, want) {
		t.Errorf("Request body = %+v, want %+v", v, want)
	}
}
