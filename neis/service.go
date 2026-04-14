package neis

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/jihoonwjj/datagsm-openapi-sdk-go/internal/transport"
)

const dateLayout = "2006-01-02"

// NEISService는 NEIS 관련 API를 정의하는 인터페이스입니다.
type NEISService interface {
	// Meals는 날짜 조건에 따라 급식 정보를 반환합니다.
	Meals(ctx context.Context, q MealQuery) ([]Meal, error)
	// Schedules는 날짜 조건에 따라 학사일정 목록을 반환합니다.
	Schedules(ctx context.Context, q ScheduleQuery) ([]Schedule, error)
	// Timetables는 학년·반·날짜 조건에 따라 시간표 목록을 반환합니다.
	// Grade와 ClassNum은 필수 파라미터입니다.
	Timetables(ctx context.Context, q TimetableQuery) ([]Timetable, error)
}

// MealQuery는 급식 정보 조회 파라미터입니다.
type MealQuery struct {
	Date     *time.Time
	FromDate *time.Time
	ToDate   *time.Time
}

// ScheduleQuery는 학사일정 조회 파라미터입니다.
type ScheduleQuery struct {
	Date     *time.Time
	FromDate *time.Time
	ToDate   *time.Time
}

// TimetableQuery는 시간표 조회 파라미터입니다.
// Grade와 ClassNum은 필수입니다.
type TimetableQuery struct {
	Grade    int
	ClassNum int
	Date     *time.Time
	FromDate *time.Time
	ToDate   *time.Time
}

// NewService는 독립적으로 사용 가능한 NEIS 서비스를 생성합니다.
func NewService(baseURL, apiKey string, httpClient *http.Client) NEISService {
	return &neisService{baseURL: baseURL, apiKey: apiKey, httpClient: httpClient}
}

type neisService struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func (s *neisService) Meals(ctx context.Context, q MealQuery) ([]Meal, error) {
	var result []Meal
	if err := transport.DoGet(ctx, s.httpClient, s.baseURL+"/v1/meals", s.apiKey, buildDateParams(q.Date, q.FromDate, q.ToDate), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *neisService) Schedules(ctx context.Context, q ScheduleQuery) ([]Schedule, error) {
	var result []Schedule
	if err := transport.DoGet(ctx, s.httpClient, s.baseURL+"/v1/schedules", s.apiKey, buildDateParams(q.Date, q.FromDate, q.ToDate), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *neisService) Timetables(ctx context.Context, q TimetableQuery) ([]Timetable, error) {
	params := buildDateParams(q.Date, q.FromDate, q.ToDate)
	params["grade"] = strconv.Itoa(q.Grade)
	params["classNum"] = strconv.Itoa(q.ClassNum)
	var result []Timetable
	if err := transport.DoGet(ctx, s.httpClient, s.baseURL+"/v1/timetables", s.apiKey, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func buildDateParams(date, fromDate, toDate *time.Time) map[string]string {
	p := make(map[string]string)
	if date != nil {
		p["date"] = date.Format(dateLayout)
	}
	if fromDate != nil {
		p["fromDate"] = fromDate.Format(dateLayout)
	}
	if toDate != nil {
		p["toDate"] = toDate.Format(dateLayout)
	}
	return p
}
