package zaico

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ErrorResponse ZAICO APIのエラーレスポンス
type ErrorResponse struct {
	Response *http.Response
	Code     int    `json:"code"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Status, r.Message)
}

// CheckResponse レスポンスのステータスコードをチェックし、エラーの場合はErrorResponseを返します
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}
