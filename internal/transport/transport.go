package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jihoonwjj/datagsm-openapi-sdk-go/internal/apierror"
)

// apiResponse는 DataGSM OpenAPI의 공통 응답 래퍼입니다.
type apiResponse[T any] struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// DoGet은 HTTP GET 요청을 실행하고 응답 data 필드를 target에 JSON 디코딩합니다.
//   - apiKey: X-API-KEY 헤더에 주입되는 인증 키
//   - rawURL: 기본 URL + 경로 (예: "https://openapi.datagsm.kr/v1/students")
//   - params: 쿼리 파라미터 맵
//   - target: 응답 data 필드를 디코딩할 대상 포인터
func DoGet(ctx context.Context, client *http.Client, rawURL string, apiKey string, params map[string]string, target any) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("datagsm: invalid URL %q: %w", rawURL, err)
	}

	if len(params) > 0 {
		q := u.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("datagsm: failed to create request: %w", err)
	}
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("datagsm: request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("datagsm: failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp struct {
			Message string `json:"message"`
		}
		msg := http.StatusText(resp.StatusCode)
		if json.Unmarshal(body, &errResp) == nil && errResp.Message != "" {
			msg = errResp.Message
		}
		return &apierror.APIError{StatusCode: resp.StatusCode, Message: msg}
	}

	var wrapper apiResponse[json.RawMessage]
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return fmt.Errorf("datagsm: failed to decode response: %w", err)
	}

	if err := json.Unmarshal(wrapper.Data, target); err != nil {
		return fmt.Errorf("datagsm: failed to decode data field: %w", err)
	}

	return nil
}
