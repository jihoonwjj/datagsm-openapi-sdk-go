package student

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jihoonwjj/datagsm-openapi-sdk-go/internal/transport"
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/types"
)

// StudentService는 학생 관련 API를 정의하는 인터페이스입니다.
type StudentService interface {
	// List는 필터·페이징·정렬 조건에 따라 학생 목록을 반환합니다.
	List(ctx context.Context, q StudentQuery) (*StudentResponse, error)
	// Get은 학생 ID로 특정 학생 정보를 반환합니다.
	Get(ctx context.Context, id int64) (*Student, error)
}

// StudentQuery는 학생 목록 조회 파라미터입니다.
type StudentQuery struct {
	Name             *string            `query:"name"`
	Email            *string            `query:"email"`
	Grade            *int               `query:"grade"`
	ClassNum         *int               `query:"classNum"`
	Number           *int               `query:"number"`
	Major            *types.Major       `query:"major"`
	Sex              *types.Sex         `query:"sex"`
	Role             *StudentRole       `query:"role"`
	StudentNumber    *int               `query:"studentNumber"`
	OnlyEnrolled     *bool              `query:"onlyEnrolled"`
	IncludeGraduates *bool              `query:"includeGraduates"`
	IncludeWithdrawn *bool              `query:"includeWithdrawn"`
	Page             int                `query:"page"`
	Size             int                `query:"size"`
	SortBy           StudentSortBy      `query:"sortBy"`
	SortDir          types.SortDirection `query:"sortDirection"`
}

// NewService는 독립적으로 사용 가능한 학생 서비스를 생성합니다.
func NewService(baseURL, apiKey string, httpClient *http.Client) StudentService {
	return &studentService{baseURL: baseURL, apiKey: apiKey, httpClient: httpClient}
}

type studentService struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func (s *studentService) List(ctx context.Context, q StudentQuery) (*StudentResponse, error) {
	var result StudentResponse
	if err := transport.DoGet(ctx, s.httpClient, s.baseURL+"/v1/students", s.apiKey, buildStudentParams(q), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *studentService) Get(ctx context.Context, id int64) (*Student, error) {
	var result Student
	url := fmt.Sprintf("%s/v1/students/%d", s.baseURL, id)
	if err := transport.DoGet(ctx, s.httpClient, url, s.apiKey, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func buildStudentParams(q StudentQuery) map[string]string {
	p := make(map[string]string)
	if q.Name != nil {
		p["name"] = *q.Name
	}
	if q.Email != nil {
		p["email"] = *q.Email
	}
	if q.Grade != nil {
		p["grade"] = strconv.Itoa(*q.Grade)
	}
	if q.ClassNum != nil {
		p["classNum"] = strconv.Itoa(*q.ClassNum)
	}
	if q.Number != nil {
		p["number"] = strconv.Itoa(*q.Number)
	}
	if q.Major != nil {
		p["major"] = string(*q.Major)
	}
	if q.Sex != nil {
		p["sex"] = string(*q.Sex)
	}
	if q.Role != nil {
		p["role"] = string(*q.Role)
	}
	if q.StudentNumber != nil {
		p["studentNumber"] = strconv.Itoa(*q.StudentNumber)
	}
	if q.OnlyEnrolled != nil {
		p["onlyEnrolled"] = strconv.FormatBool(*q.OnlyEnrolled)
	}
	if q.IncludeGraduates != nil {
		p["includeGraduates"] = strconv.FormatBool(*q.IncludeGraduates)
	}
	if q.IncludeWithdrawn != nil {
		p["includeWithdrawn"] = strconv.FormatBool(*q.IncludeWithdrawn)
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
