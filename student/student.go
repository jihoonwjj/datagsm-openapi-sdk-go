// Package student는 DataGSM 학생 도메인 타입과 API를 제공합니다.
package student

import (
	"github.com/themoment-team/datagsm-openapi-sdk-go/club"
	"github.com/themoment-team/datagsm-openapi-sdk-go/types"
)

// StudentRole은 학생의 역할을 나타냅니다.
type StudentRole string

const (
	StudentRoleGeneralStudent   StudentRole = "GENERAL_STUDENT"
	StudentRoleStudentCouncil   StudentRole = "STUDENT_COUNCIL"
	StudentRoleDormitoryManager StudentRole = "DORMITORY_MANAGER"
	StudentRoleGraduate         StudentRole = "GRADUATE"
	StudentRoleWithdrawn        StudentRole = "WITHDRAWN"
)

// StudentSortBy는 학생 목록 정렬 기준을 나타냅니다.
type StudentSortBy string

const (
	StudentSortByID            StudentSortBy = "ID"
	StudentSortByName          StudentSortBy = "NAME"
	StudentSortByEmail         StudentSortBy = "EMAIL"
	StudentSortByStudentNumber StudentSortBy = "STUDENT_NUMBER"
	StudentSortByGrade         StudentSortBy = "GRADE"
	StudentSortByClassNum      StudentSortBy = "CLASS_NUM"
	StudentSortByNumber        StudentSortBy = "NUMBER"
	StudentSortByMajor         StudentSortBy = "MAJOR"
	StudentSortByRole          StudentSortBy = "ROLE"
	StudentSortBySex           StudentSortBy = "SEX"
	StudentSortByDormitoryRoom StudentSortBy = "DORMITORY_ROOM"
)

// Student는 학생 정보를 나타냅니다.
type Student struct {
	ID             int64       `json:"id"`
	Name           string      `json:"name"`
	Sex            types.Sex   `json:"sex"`
	Email          string      `json:"email"`
	Grade          int         `json:"grade"`
	ClassNum       int         `json:"classNum"`
	Number         int         `json:"number"`
	StudentNumber  int         `json:"studentNumber"`
	Major          types.Major `json:"major"`
	Specialty      string      `json:"specialty"`
	Role           StudentRole `json:"role"`
	DormitoryFloor *int        `json:"dormitoryFloor,omitempty"`
	DormitoryRoom  *int        `json:"dormitoryRoom,omitempty"`
	MajorClub      *club.Club  `json:"majorClub,omitempty"`
	AutonomousClub *club.Club  `json:"autonomousClub,omitempty"`
	GithubID       *string     `json:"githubId,omitempty"`
	GithubURL      *string     `json:"githubUrl,omitempty"`
}

// StudentResponse는 학생 목록 조회 결과입니다.
type StudentResponse struct {
	Students      []Student `json:"students"`
	TotalElements int64     `json:"totalElements"`
	TotalPages    int       `json:"totalPages"`
}
