// Package club은 DataGSM 동아리 도메인 타입과 API를 제공합니다.
package club

import "github.com/themoment-team/datagsm-openapi-sdk-go/types"

// ClubType은 동아리 유형을 나타냅니다.
type ClubType string

const (
	ClubTypeMajor      ClubType = "MAJOR_CLUB"
	ClubTypeAutonomous ClubType = "AUTONOMOUS_CLUB"
)

// ClubStatus는 동아리 상태를 나타냅니다.
type ClubStatus string

const (
	ClubStatusActive    ClubStatus = "ACTIVE"
	ClubStatusAbolished ClubStatus = "ABOLISHED"
)

// ClubSortBy는 동아리 목록 정렬 기준을 나타냅니다.
type ClubSortBy string

const (
	ClubSortByID            ClubSortBy = "ID"
	ClubSortByName          ClubSortBy = "NAME"
	ClubSortByType          ClubSortBy = "TYPE"
	ClubSortByFoundedYear   ClubSortBy = "FOUNDED_YEAR"
	ClubSortByAbolishedYear ClubSortBy = "ABOLISHED_YEAR"
	ClubSortByStatus        ClubSortBy = "STATUS"
)

// Club은 동아리 기본 정보를 나타냅니다.
type Club struct {
	ID            int64      `json:"id"`
	Name          string     `json:"name"`
	Type          ClubType   `json:"type"`
	Status        ClubStatus `json:"status"`
	FoundedYear   int        `json:"foundedYear"`
	AbolishedYear *int       `json:"abolishedYear,omitempty"`
}

// ParticipantInfo는 동아리 또는 프로젝트 참여자 정보를 나타냅니다.
type ParticipantInfo struct {
	ID            int64        `json:"id"`
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	StudentNumber *int         `json:"studentNumber,omitempty"`
	Major         *types.Major `json:"major,omitempty"`
	Sex           types.Sex    `json:"sex"`
}

// ClubDetail은 동아리 상세 정보(리더·참여자 포함)를 나타냅니다.
type ClubDetail struct {
	Club
	Leader       *ParticipantInfo  `json:"leader,omitempty"`
	Participants []ParticipantInfo `json:"participants"`
}

// ClubResponse는 동아리 목록 조회 결과입니다.
type ClubResponse struct {
	Clubs         []ClubDetail `json:"clubs"`
	TotalElements int64        `json:"totalElements"`
	TotalPages    int          `json:"totalPages"`
}
