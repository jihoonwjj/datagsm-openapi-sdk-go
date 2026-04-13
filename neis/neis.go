// Package neis는 DataGSM NEIS(급식·학사일정·시간표) 도메인 타입과 API를 제공합니다.
package neis

import "time"

// MealType은 급식 유형을 나타냅니다.
type MealType string

const (
	MealTypeBreakfast MealType = "BREAKFAST"
	MealTypeLunch     MealType = "LUNCH"
	MealTypeDinner    MealType = "DINNER"
)

// Meal은 급식 정보를 나타냅니다.
type Meal struct {
	MealID          string    `json:"mealId"`
	SchoolCode      string    `json:"schoolCode"`
	SchoolName      string    `json:"schoolName"`
	OfficeCode      string    `json:"officeCode"`
	OfficeName      string    `json:"officeName"`
	MealDate        time.Time `json:"mealDate"`
	MealType        MealType  `json:"mealType"`
	MealMenu        []string  `json:"mealMenu"`
	MealAllergyInfo []string  `json:"mealAllergyInfo,omitempty"`
	MealCalories    *string   `json:"mealCalories,omitempty"`
	OriginInfo      *string   `json:"originInfo,omitempty"`
	NutritionInfo   *string   `json:"nutritionInfo,omitempty"`
	MealServeCount  *int      `json:"mealServeCount,omitempty"`
}

// Schedule은 학사일정 정보를 나타냅니다.
type Schedule struct {
	ScheduleID       string    `json:"scheduleId"`
	SchoolCode       string    `json:"schoolCode"`
	SchoolName       string    `json:"schoolName"`
	OfficeCode       string    `json:"officeCode"`
	OfficeName       string    `json:"officeName"`
	ScheduleDate     time.Time `json:"scheduleDate"`
	AcademicYear     int       `json:"academicYear"`
	EventName        string    `json:"eventName"`
	EventContent     *string   `json:"eventContent,omitempty"`
	DayCategory      *string   `json:"dayCategory,omitempty"`
	SchoolCourseType *string   `json:"schoolCourseType,omitempty"`
	DayNightType     *string   `json:"dayNightType,omitempty"`
	TargetGrades     []int     `json:"targetGrades"`
}

// Timetable은 시간표 정보를 나타냅니다.
type Timetable struct {
	TimetableID   string    `json:"timetableId"`
	SchoolCode    string    `json:"schoolCode"`
	SchoolName    string    `json:"schoolName"`
	OfficeCode    string    `json:"officeCode"`
	OfficeName    string    `json:"officeName"`
	TimetableDate time.Time `json:"timetableDate"`
	AcademicYear  int       `json:"academicYear"`
	Semester      int       `json:"semester"`
	Grade         int       `json:"grade"`
	ClassNum      int       `json:"classNum"`
	Period        int       `json:"period"`
	Subject       *string   `json:"subject,omitempty"`
}
