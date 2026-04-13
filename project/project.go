// Package project는 DataGSM 프로젝트 도메인 타입과 API를 제공합니다.
package project

import "github.com/themoment-team/datagsm-openapi-sdk-go/club"

// ProjectSortBy는 프로젝트 목록 정렬 기준을 나타냅니다.
type ProjectSortBy string

const (
	ProjectSortByID   ProjectSortBy = "ID"
	ProjectSortByName ProjectSortBy = "NAME"
)

// Project는 프로젝트 정보를 나타냅니다.
type Project struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Club         *club.Club         `json:"club,omitempty"`
	Participants []club.ParticipantInfo `json:"participants"`
}

// ProjectResponse는 프로젝트 목록 조회 결과입니다.
type ProjectResponse struct {
	Projects      []Project `json:"projects"`
	TotalElements int64     `json:"totalElements"`
	TotalPages    int       `json:"totalPages"`
}
