// Package apierror는 DataGSM OpenAPI HTTP 오류 타입을 정의합니다.
// transport와 서브 패키지가 공유하는 단일 에러 타입입니다.
package apierror

import "fmt"

// APIError는 DataGSM OpenAPI 서버에서 반환된 HTTP 오류를 나타냅니다.
// errors.As를 사용하여 상태 코드를 확인하세요.
//
//	var apiErr *datagsm.APIError
//	if errors.As(err, &apiErr) {
//	    fmt.Println(apiErr.StatusCode) // e.g. 401
//	}
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("datagsm: API error %d: %s", e.StatusCode, e.Message)
}
