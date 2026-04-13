package datagsm

import (
	"net/http"
	"time"
)

// ClientOption은 DataGsmClient를 설정하는 함수형 옵션입니다.
type ClientOption func(*DataGsmClient)

// WithBaseURL은 API 기본 URL을 설정합니다. (기본값: "https://openapi.datagsm.kr")
func WithBaseURL(url string) ClientOption {
	return func(c *DataGsmClient) {
		c.baseURL = url
	}
}

// WithTimeout은 HTTP 요청의 타임아웃을 설정합니다. (기본값: 30초)
func WithTimeout(d time.Duration) ClientOption {
	return func(c *DataGsmClient) {
		c.httpClient.Timeout = d
	}
}

// WithHTTPClient는 사용자 정의 *http.Client를 주입합니다.
// 이 옵션을 사용하면 WithTimeout 설정은 무시됩니다.
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *DataGsmClient) {
		c.httpClient = hc
	}
}
