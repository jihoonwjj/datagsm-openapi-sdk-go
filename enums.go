package datagsm

// 이 파일은 서브 패키지에 정의된 열거형 상수를 루트 패키지에서
// 그대로 참조할 수 있도록 재노출합니다.

import (
	"github.com/themoment-team/datagsm-openapi-sdk-go/club"
	"github.com/themoment-team/datagsm-openapi-sdk-go/neis"
	"github.com/themoment-team/datagsm-openapi-sdk-go/student"
	"github.com/themoment-team/datagsm-openapi-sdk-go/types"
)

// ── types 패키지 상수 ────────────────────────────────────────────────────────

const (
	MajorSWDevelopment = types.MajorSWDevelopment
	MajorSmartIoT      = types.MajorSmartIoT
	MajorAI            = types.MajorAI

	SexMan   = types.SexMan
	SexWoman = types.SexWoman

	SortDirASC  = types.SortDirASC
	SortDirDESC = types.SortDirDESC
)

// ── club 패키지 상수 ─────────────────────────────────────────────────────────

const (
	ClubTypeMajor      = club.ClubTypeMajor
	ClubTypeAutonomous = club.ClubTypeAutonomous

	ClubStatusActive    = club.ClubStatusActive
	ClubStatusAbolished = club.ClubStatusAbolished

	ClubSortByID            = club.ClubSortByID
	ClubSortByName          = club.ClubSortByName
	ClubSortByType          = club.ClubSortByType
	ClubSortByFoundedYear   = club.ClubSortByFoundedYear
	ClubSortByAbolishedYear = club.ClubSortByAbolishedYear
	ClubSortByStatus        = club.ClubSortByStatus
)

// ── student 패키지 상수 ──────────────────────────────────────────────────────

const (
	StudentRoleGeneralStudent   = student.StudentRoleGeneralStudent
	StudentRoleStudentCouncil   = student.StudentRoleStudentCouncil
	StudentRoleDormitoryManager = student.StudentRoleDormitoryManager
	StudentRoleGraduate         = student.StudentRoleGraduate
	StudentRoleWithdrawn        = student.StudentRoleWithdrawn

	StudentSortByID            = student.StudentSortByID
	StudentSortByName          = student.StudentSortByName
	StudentSortByEmail         = student.StudentSortByEmail
	StudentSortByStudentNumber = student.StudentSortByStudentNumber
	StudentSortByGrade         = student.StudentSortByGrade
	StudentSortByClassNum      = student.StudentSortByClassNum
	StudentSortByNumber        = student.StudentSortByNumber
	StudentSortByMajor         = student.StudentSortByMajor
	StudentSortByRole          = student.StudentSortByRole
	StudentSortBySex           = student.StudentSortBySex
	StudentSortByDormitoryRoom = student.StudentSortByDormitoryRoom
)

// ── neis 패키지 상수 ─────────────────────────────────────────────────────────

const (
	MealTypeBreakfast = neis.MealTypeBreakfast
	MealTypeLunch     = neis.MealTypeLunch
	MealTypeDinner    = neis.MealTypeDinner
)
