package club

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/themoment-team/datagsm-openapi-sdk-go/internal/transport"
	"github.com/themoment-team/datagsm-openapi-sdk-go/types"
)

// ClubService는 동아리 관련 API를 정의하는 인터페이스입니다.
type ClubService interface {
	// List는 필터·페이징·정렬 조건에 따라 동아리 목록을 반환합니다.
	List(ctx context.Context, q ClubQuery) (*ClubResponse, error)
	// Get은 동아리 ID로 특정 동아리의 상세 정보를 반환합니다.
	Get(ctx context.Context, id int64) (*ClubDetail, error)
}

// ClubQuery는 동아리 목록 조회 파라미터입니다.
type ClubQuery struct {
	ClubID                      *int64             `query:"clubId"`
	ClubName                    *string            `query:"clubName"`
	ClubType                    *ClubType          `query:"clubType"`
	Status                      *ClubStatus        `query:"status"`
	FoundedYear                 *int               `query:"foundedYear"`
	IncludeLeaderInParticipants *bool              `query:"includeLeaderInParticipants"`
	Page                        int                `query:"page"`
	Size                        int                `query:"size"`
	SortBy                      ClubSortBy         `query:"sortBy"`
	SortDir                     types.SortDirection `query:"sortDirection"`
}

// NewService는 독립적으로 사용 가능한 동아리 서비스를 생성합니다.
func NewService(baseURL, apiKey string, httpClient *http.Client) ClubService {
	return &clubService{baseURL: baseURL, apiKey: apiKey, httpClient: httpClient}
}

type clubService struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func (s *clubService) List(ctx context.Context, q ClubQuery) (*ClubResponse, error) {
	var result ClubResponse
	if err := transport.DoGet(ctx, s.httpClient, s.baseURL+"/v1/clubs", s.apiKey, buildClubParams(q), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *clubService) Get(ctx context.Context, id int64) (*ClubDetail, error) {
	var result ClubDetail
	url := fmt.Sprintf("%s/v1/clubs/%d", s.baseURL, id)
	if err := transport.DoGet(ctx, s.httpClient, url, s.apiKey, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func buildClubParams(q ClubQuery) map[string]string {
	p := make(map[string]string)
	if q.ClubID != nil {
		p["clubId"] = strconv.FormatInt(*q.ClubID, 10)
	}
	if q.ClubName != nil {
		p["clubName"] = *q.ClubName
	}
	if q.ClubType != nil {
		p["clubType"] = string(*q.ClubType)
	}
	if q.Status != nil {
		p["status"] = string(*q.Status)
	}
	if q.FoundedYear != nil {
		p["foundedYear"] = strconv.Itoa(*q.FoundedYear)
	}
	if q.IncludeLeaderInParticipants != nil {
		p["includeLeaderInParticipants"] = strconv.FormatBool(*q.IncludeLeaderInParticipants)
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
