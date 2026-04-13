// Package types는 DataGSM SDK 여러 도메인에서 공유하는 기본 타입을 제공합니다.
package types

// Major는 학생의 전공을 나타냅니다.
type Major string

const (
	MajorSWDevelopment Major = "SW_DEVELOPMENT"
	MajorSmartIoT      Major = "SMART_IOT"
	MajorAI            Major = "AI"
)

// Sex는 성별을 나타냅니다.
type Sex string

const (
	SexMan   Sex = "MAN"
	SexWoman Sex = "WOMAN"
)

// SortDirection은 정렬 방향을 나타냅니다.
type SortDirection string

const (
	SortDirASC  SortDirection = "ASC"
	SortDirDESC SortDirection = "DESC"
)
