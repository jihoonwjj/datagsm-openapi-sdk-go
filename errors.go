package datagsm

import "github.com/themoment-team/datagsm-openapi-sdk-go/internal/apierror"

// APIError는 DataGSM OpenAPI 서버에서 반환된 HTTP 오류를 나타냅니다.
// errors.As를 사용하여 상태 코드를 확인하세요.
//
//	var apiErr *datagsm.APIError
//	if errors.As(err, &apiErr) {
//	    switch apiErr.StatusCode {
//	    case 401: // 인증 오류
//	    case 429: // 속도 제한
//	    }
//	}
type APIError = apierror.APIError
