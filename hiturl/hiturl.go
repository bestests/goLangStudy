package hiturl

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFailed = errors.New("ERROR - Request Failed")
)

// HitURL - check url
func HitURL(url string) error {

	fmt.Println("CHECK URL - ", url)

	resp, err := http.Get(url)

	respStatusCode := resp.StatusCode

	result := "OK"

	if err != nil || respStatusCode >= 400 {
		result = "FAILED"
		fmt.Printf("ERROR CODE - %-3d : %-8s\n", respStatusCode, result)
		return errRequestFailed
	}

	fmt.Printf("URL : %-28s - %-3d - %-8s\n", url, respStatusCode, result)

	return nil
}
