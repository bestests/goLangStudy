package hiturl

import (
	"fmt"
	"net/http"
)

// ResponseResult - HitURL result struct
type ResponseResult struct {
	url    string
	status string
	code   int
}

// HitURL2 - use goroutine hitURL
func HitURL2(url string, c chan ResponseResult) {
	resp, err := http.Get(url)

	status := "OK"

	code := resp.StatusCode

	if err != nil || code >= 400 {
		status = "FAILED"
	}

	c <- ResponseResult{url: url, status: status, code: code}
}

func (r ResponseResult) String() string {
	return fmt.Sprintf("URL : %-31s | STATUS : %-6s | CODE : %3d", r.url, r.status, r.code)
}
