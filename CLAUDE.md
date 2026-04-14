# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# 전체 테스트 실행
go test ./...

# 단일 테스트 실행
go test -run TestNewClient_ValidAPIKey .

# 특정 패키지 테스트
go test ./student/...

# 빌드 검증
go build ./...

# 정적 분석
go vet ./...

# 예제 실행 (API 키 필요)
DATAGSM_API_KEY=your-key go run examples/main.go
```

## Architecture

이 SDK는 **네임스페이스 서비스 패턴**을 사용합니다. 루트 패키지(`datagsm`)가 진입점이며, 도메인별 서브패키지에 실제 구현이 있습니다.

### 패키지 구조

```
datagsm (루트 패키지)
├── client.go       — DataGsmClient, NewClient, Close
├── options.go      — ClientOption (WithBaseURL, WithTimeout, WithHTTPClient)
├── aliases.go      — 서브패키지 타입을 루트에 재노출 (datagsm.Student 등 사용 가능)
├── errors.go       — APIError (internal/apierror의 앨리어스)
├── enums.go        — 공통 열거형 상수
├── common.go       — Ptr[T] 제네릭 헬퍼
│
├── student/        — StudentService 인터페이스 + 구현 + 타입
├── club/           — ClubService 인터페이스 + 구현 + 타입
├── project/        — ProjectService 인터페이스 + 구현 + 타입
├── neis/           — NEISService 인터페이스 + 구현 + 타입
├── types/          — 공유 원시 타입 (Major, Sex, SortDirection)
│
└── internal/
    ├── apierror/   — APIError 타입 정의
    └── transport/  — DoGet: HTTP GET 실행, 응답 언래핑, 에러 처리
```

### 데이터 흐름

1. 호출자: `client.Students().List(ctx, query)` (루트 패키지 타입 앨리어스 사용)
2. `student.studentService.List()`: 쿼리 파라미터 빌드 → `transport.DoGet()` 호출
3. `transport.DoGet()`: X-API-KEY 헤더 주입 → HTTP GET → `apiResponse[T]` 언래핑 → JSON 디코딩
4. 비 2xx 응답 시 `apierror.APIError{StatusCode, Message}` 반환

### 새 서비스 추가 패턴

각 도메인 패키지(`student/`, `club/` 등)는 동일한 구조를 따릅니다:
- `types.go` (또는 `{domain}.go`): 도메인 타입 정의
- `service.go`: Service 인터페이스 + 구현체 + 쿼리 파라미터 빌더

새 서비스를 추가하면 `client.go`에 필드와 메서드를 추가하고, `aliases.go`에 타입 앨리어스를 선언해야 합니다.

## Key Design Decisions

- **인증**: `X-API-KEY` 헤더만 사용 (Bearer 토큰 없음)
- **HTTP 클라이언트**: 표준 `net/http` (외부 의존성 없음)
- **API 응답 형식**: 모든 응답은 `{status, code, message, data}` 래퍼로 감싸져 있음; `transport.DoGet`이 `data` 필드만 추출
- **선택적 파라미터**: 포인터 필드 사용 + `Ptr[T]` 헬퍼로 인라인 생성
- **Base URL**: `https://openapi.datagsm.kr` (기본값)
