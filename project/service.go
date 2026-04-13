package project

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/themoment-team/datagsm-openapi-sdk-go/internal/transport"
	"github.com/themoment-team/datagsm-openapi-sdk-go/types"
)

// ProjectService는 프로젝트 관련 API를 정의하는 인터페이스입니다.
type ProjectService interface {
	// List는 필터·페이징·정렬 조건에 따라 프로젝트 목록을 반환합니다.
	List(ctx context.Context, q ProjectQuery) (*ProjectResponse, error)
	// Get은 프로젝트 ID로 특정 프로젝트 정보를 반환합니다.
	Get(ctx context.Context, id int64) (*Project, error)
}

// ProjectQuery는 프로젝트 목록 조회 파라미터입니다.
type ProjectQuery struct {
	ProjectID   *int64             `query:"projectId"`
	ProjectName *string            `query:"projectName"`
	ClubID      *int64             `query:"clubId"`
	Page        int                `query:"page"`
	Size        int                `query:"size"`
	SortBy      ProjectSortBy      `query:"sortBy"`
	SortDir     types.SortDirection `query:"sortDirection"`
}

// NewService는 독립적으로 사용 가능한 프로젝트 서비스를 생성합니다.
func NewService(baseURL, apiKey string, httpClient *http.Client) ProjectService {
	return &projectService{baseURL: baseURL, apiKey: apiKey, httpClient: httpClient}
}

type projectService struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func (s *projectService) List(ctx context.Context, q ProjectQuery) (*ProjectResponse, error) {
	var result ProjectResponse
	if err := transport.DoGet(ctx, s.httpClient, s.baseURL+"/v1/projects", s.apiKey, buildProjectParams(q), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *projectService) Get(ctx context.Context, id int64) (*Project, error) {
	var result Project
	url := fmt.Sprintf("%s/v1/projects/%d", s.baseURL, id)
	if err := transport.DoGet(ctx, s.httpClient, url, s.apiKey, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func buildProjectParams(q ProjectQuery) map[string]string {
	p := make(map[string]string)
	if q.ProjectID != nil {
		p["projectId"] = strconv.FormatInt(*q.ProjectID, 10)
	}
	if q.ProjectName != nil {
		p["projectName"] = *q.ProjectName
	}
	if q.ClubID != nil {
		p["clubId"] = strconv.FormatInt(*q.ClubID, 10)
	}
	p["page"] = strconv.Itoa(q.Page)
	if q.Size > 0 {
		p["size"] = strconv.Itoa(q.Size)
	}
	if q.SortBy != "" {
		p["sortBy"] = string(q.SortBy)
	}
	if q.SortDir != "" {
		p["sortDirection"] = string(q.SortDir)
	}
	return p
}
