package datagsm_test

import (
	"context"
	"net/http"
	"testing"

	datagsm "github.com/themoment-team/datagsm-openapi-sdk-go"
)

func TestProjects_List(t *testing.T) {
	want := datagsm.ProjectResponse{
		Projects: []datagsm.Project{
			{ID: 1, Name: "DataGSM", Description: "학교 데이터 플랫폼", Participants: []datagsm.ParticipantInfo{}},
		},
		TotalElements: 1,
		TotalPages:    1,
	}

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/projects" {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, want)
	}))

	resp, err := client.Projects().List(context.Background(), datagsm.ProjectQuery{Page: 0, Size: 10})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(resp.Projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(resp.Projects))
	}
	if resp.Projects[0].Name != "DataGSM" {
		t.Errorf("expected name 'DataGSM', got %q", resp.Projects[0].Name)
	}
}

func TestProjects_Get(t *testing.T) {
	want := datagsm.Project{ID: 10, Name: "오픈API", Description: "공개 API 프로젝트", Participants: []datagsm.ParticipantInfo{}}

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/projects/10" {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, want)
	}))

	got, err := client.Projects().Get(context.Background(), 10)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got.ID != 10 {
		t.Errorf("expected ID 10, got %d", got.ID)
	}
}
