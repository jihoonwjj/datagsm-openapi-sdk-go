// examples/main.go는 DataGSM OpenAPI Go SDK의 전체 사용 예시입니다.
//
// 실행 방법:
//
//	DATAGSM_API_KEY=your-api-key go run examples/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	datagsm "github.com/jihoonwjj/datagsm-openapi-sdk-go"
)

func main() {
	apiKey := os.Getenv("DATAGSM_API_KEY")
	if apiKey == "" {
		log.Fatal("DATAGSM_API_KEY 환경변수가 설정되지 않았습니다")
	}

	client, err := datagsm.NewClient(apiKey,
		datagsm.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatalf("클라이언트 생성 실패: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// ─── 학생 목록 조회 ───────────────────────────────────────────
	fmt.Println("=== 학생 목록 (1학년, 소프트웨어개발과) ===")
	students, err := client.Students().List(ctx, datagsm.StudentQuery{
		Grade: datagsm.Ptr(1),
		Major: (*datagsm.Major)(datagsm.Ptr(string(datagsm.MajorSWDevelopment))),
		Page:  0,
		Size:  5,
	})
	if err != nil {
		printAPIError("Students.List", err)
	} else {
		fmt.Printf("총 %d명 (페이지: %d)\n", students.TotalElements, students.TotalPages)
		for _, s := range students.Students {
			fmt.Printf("  [%d] %s — %d학년 %d반 %d번\n", s.ID, s.Name, s.Grade, s.ClassNum, s.Number)
		}
	}

	// ─── 동아리 목록 조회 ──────────────────────────────────────────
	fmt.Println("\n=== 활동 중인 전공동아리 목록 ===")
	clubs, err := client.Clubs().List(ctx, datagsm.ClubQuery{
		ClubType: (*datagsm.ClubType)(datagsm.Ptr(string(datagsm.ClubTypeMajor))),
		Status:   (*datagsm.ClubStatus)(datagsm.Ptr(string(datagsm.ClubStatusActive))),
		Page:     0,
		Size:     10,
	})
	if err != nil {
		printAPIError("Clubs.List", err)
	} else {
		fmt.Printf("총 %d개\n", clubs.TotalElements)
		for _, c := range clubs.Clubs {
			fmt.Printf("  [%d] %s (%d년 창설)\n", c.ID, c.Name, c.FoundedYear)
		}
	}

	// ─── 프로젝트 목록 조회 ────────────────────────────────────────
	fmt.Println("\n=== 프로젝트 목록 ===")
	projects, err := client.Projects().List(ctx, datagsm.ProjectQuery{
		Page: 0,
		Size: 5,
	})
	if err != nil {
		printAPIError("Projects.List", err)
	} else {
		fmt.Printf("총 %d개\n", projects.TotalElements)
		for _, p := range projects.Projects {
			fmt.Printf("  [%d] %s\n", p.ID, p.Name)
		}
	}

	// ─── 오늘 급식 조회 ────────────────────────────────────────────
	fmt.Println("\n=== 오늘 급식 ===")
	today := time.Now()
	meals, err := client.NEIS().Meals(ctx, datagsm.MealQuery{Date: &today})
	if err != nil {
		printAPIError("NEIS.Meals", err)
	} else {
		for _, m := range meals {
			fmt.Printf("  [%s] %s\n", m.MealType, joinStrings(m.MealMenu))
		}
		if len(meals) == 0 {
			fmt.Println("  (오늘 급식 정보 없음)")
		}
	}

	// ─── 이번 달 학사일정 조회 ──────────────────────────────────────
	fmt.Println("\n=== 이번 달 학사일정 ===")
	from := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local)
	to := from.AddDate(0, 1, -1)
	schedules, err := client.NEIS().Schedules(ctx, datagsm.ScheduleQuery{
		FromDate: &from,
		ToDate:   &to,
	})
	if err != nil {
		printAPIError("NEIS.Schedules", err)
	} else {
		for _, s := range schedules {
			fmt.Printf("  %s — %s\n", s.ScheduleDate.Format("01/02"), s.EventName)
		}
		if len(schedules) == 0 {
			fmt.Println("  (이번 달 학사일정 없음)")
		}
	}

	// ─── 시간표 조회 (1학년 1반) ────────────────────────────────────
	fmt.Println("\n=== 오늘 1학년 1반 시간표 ===")
	timetables, err := client.NEIS().Timetables(ctx, datagsm.TimetableQuery{
		Grade:    1,
		ClassNum: 1,
		Date:     &today,
	})
	if err != nil {
		printAPIError("NEIS.Timetables", err)
	} else {
		for _, tt := range timetables {
			subject := "(없음)"
			if tt.Subject != nil {
				subject = *tt.Subject
			}
			fmt.Printf("  %d교시: %s\n", tt.Period, subject)
		}
		if len(timetables) == 0 {
			fmt.Println("  (시간표 정보 없음)")
		}
	}
}

func printAPIError(op string, err error) {
	fmt.Printf("  [오류] %s: %v\n", op, err)
}

func joinStrings(ss []string) string {
	result := ""
	for i, s := range ss {
		if i > 0 {
			result += ", "
		}
		result += s
	}
	return result
}
