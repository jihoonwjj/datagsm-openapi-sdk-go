package datagsm

// 이 파일은 서브 패키지 타입을 루트 패키지에서 그대로 사용할 수 있도록
// 타입 앨리어스를 선언합니다. 기존 datagsm.Student, datagsm.ClubQuery 등의
// 사용 방식이 그대로 유지됩니다.

import (
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/club"
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/neis"
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/project"
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/student"
	"github.com/jihoonwjj/datagsm-openapi-sdk-go/types"
)

// ── 공유 원시 타입 ────────────────────────────────────────────────────────────

type Major         = types.Major
type Sex           = types.Sex
type SortDirection = types.SortDirection

// ── 동아리 도메인 ────────────────────────────────────────────────────────────

type Club            = club.Club
type ClubDetail      = club.ClubDetail
type ClubResponse    = club.ClubResponse
type ClubQuery       = club.ClubQuery
type ParticipantInfo = club.ParticipantInfo
type ClubType        = club.ClubType
type ClubStatus      = club.ClubStatus
type ClubSortBy      = club.ClubSortBy
type ClubService     = club.ClubService

// ── 학생 도메인 ──────────────────────────────────────────────────────────────

type Student         = student.Student
type StudentResponse = student.StudentResponse
type StudentQuery    = student.StudentQuery
type StudentRole     = student.StudentRole
type StudentSortBy   = student.StudentSortBy
type StudentService  = student.StudentService

// ── 프로젝트 도메인 ──────────────────────────────────────────────────────────

type Project         = project.Project
type ProjectResponse = project.ProjectResponse
type ProjectQuery    = project.ProjectQuery
type ProjectSortBy   = project.ProjectSortBy
type ProjectService  = project.ProjectService

// ── NEIS 도메인 ──────────────────────────────────────────────────────────────

type Meal           = neis.Meal
type Schedule       = neis.Schedule
type Timetable      = neis.Timetable
type MealQuery      = neis.MealQuery
type ScheduleQuery  = neis.ScheduleQuery
type TimetableQuery = neis.TimetableQuery
type MealType       = neis.MealType
type NEISService    = neis.NEISService
