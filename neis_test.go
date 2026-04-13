package datagsm_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	datagsm "github.com/themoment-team/datagsm-openapi-sdk-go"
)

func TestNEIS_Meals(t *testing.T) {
	want := []datagsm.Meal{
		{
			MealID:     "1",
			SchoolName: "광주소프트웨어마이스터고등학교",
			MealDate:   time.Date(2026, 4, 13, 0, 0, 0, 0, time.UTC),
			MealType:   datagsm.MealTypeLunch,
			MealMenu:   []string{"밥", "김치", "된장국"},
		},
	}

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/meals" {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, want)
	}))

	today := time.Date(2026, 4, 13, 0, 0, 0, 0, time.UTC)
	meals, err := client.NEIS().Meals(context.Background(), datagsm.MealQuery{Date: &today})
	if err != nil {
		t.Fatalf("Meals: %v", err)
	}
	if len(meals) != 1 {
		t.Fatalf("expected 1 meal, got %d", len(meals))
	}
	if meals[0].MealType != datagsm.MealTypeLunch {
		t.Errorf("expected LUNCH, got %q", meals[0].MealType)
	}
}

func TestNEIS_Timetables_QueryParams(t *testing.T) {
	var gotGrade, gotClassNum string

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotGrade = r.URL.Query().Get("grade")
		gotClassNum = r.URL.Query().Get("classNum")
		writeJSON(w, []datagsm.Timetable{})
	}))

	_, err := client.NEIS().Timetables(context.Background(), datagsm.TimetableQuery{
		Grade:    2,
		ClassNum: 3,
	})
	if err != nil {
		t.Fatalf("Timetables: %v", err)
	}
	if gotGrade != "2" {
		t.Errorf("expected grade=2, got %q", gotGrade)
	}
	if gotClassNum != "3" {
		t.Errorf("expected classNum=3, got %q", gotClassNum)
	}
}

func TestNEIS_Schedules(t *testing.T) {
	want := []datagsm.Schedule{
		{
			ScheduleID:   "s1",
			SchoolName:   "광주소프트웨어마이스터고등학교",
			ScheduleDate: time.Date(2026, 4, 13, 0, 0, 0, 0, time.UTC),
			AcademicYear: 2026,
			EventName:    "개학",
			TargetGrades: []int{1, 2, 3},
		},
	}

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/schedules" {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, want)
	}))

	schedules, err := client.NEIS().Schedules(context.Background(), datagsm.ScheduleQuery{})
	if err != nil {
		t.Fatalf("Schedules: %v", err)
	}
	if len(schedules) != 1 {
		t.Fatalf("expected 1 schedule, got %d", len(schedules))
	}
	if schedules[0].EventName != "개학" {
		t.Errorf("expected '개학', got %q", schedules[0].EventName)
	}
}
