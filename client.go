// Package datagsm는 DataGSM OpenAPI를 위한 Go SDK입니다.
//
// 기본 사용법:
//
//	client, err := datagsm.NewClient("your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer client.Close()
//
//	resp, err := client.Students().List(ctx, datagsm.StudentQuery{Page: 0, Size: 20})
package datagsm

import (
	"errors"
	"net/http"
	"time"

	"github.com/jihoonwjj/datagsm-openapi-sdk-go/club"
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/neis"
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/project"
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/student"
)

const defaultBaseURL = "https://openapi.datagsm.kr"
const defaultTimeout = 30 * time.Second

// DataGsmClient는 DataGSM OpenAPI에 접근하는 메인 클라이언트입니다.
// NewClient로 생성하고, 사용 후 Close를 호출하세요.
type DataGsmClient struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string

	students student.StudentService
	clubs    club.ClubService
	projects project.ProjectService
	neisAPI  neis.NEISService
}

// NewClient는 새 DataGsmClient를 생성합니다.
// apiKey는 필수이며 빈 문자열이면 에러를 반환합니다.
//
//	client, err := datagsm.NewClient("your-api-key",
//	    datagsm.WithBaseURL("https://openapi.datagsm.kr"),
//	    datagsm.WithTimeout(30*time.Second),
//	)
func NewClient(apiKey string, opts ...ClientOption) (*DataGsmClient, error) {
	if apiKey == "" {
		return nil, errors.New("datagsm: apiKey must not be empty")
	}

	c := &DataGsmClient{
		httpClient: &http.Client{Timeout: defaultTimeout},
		baseURL:    defaultBaseURL,
		apiKey:     apiKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.students = student.NewService(c.baseURL, c.apiKey, c.httpClient)
	c.clubs = club.NewService(c.baseURL, c.apiKey, c.httpClient)
	c.projects = project.NewService(c.baseURL, c.apiKey, c.httpClient)
	c.neisAPI = neis.NewService(c.baseURL, c.apiKey, c.httpClient)

	return c, nil
}

// Students는 학생 관련 API에 접근하는 StudentService를 반환합니다.
func (c *DataGsmClient) Students() student.StudentService {
	return c.students
}

// Clubs는 동아리 관련 API에 접근하는 ClubService를 반환합니다.
func (c *DataGsmClient) Clubs() club.ClubService {
	return c.clubs
}

// Projects는 프로젝트 관련 API에 접근하는 ProjectService를 반환합니다.
func (c *DataGsmClient) Projects() project.ProjectService {
	return c.projects
}

// NEIS는 급식·학사일정·시간표 관련 API에 접근하는 NEISService를 반환합니다.
func (c *DataGsmClient) NEIS() neis.NEISService {
	return c.neisAPI
}

// Close는 클라이언트의 유휴 연결을 정리합니다.
// defer client.Close() 패턴으로 사용하세요.
func (c *DataGsmClient) Close() error {
	c.httpClient.CloseIdleConnections()
	return nil
}
