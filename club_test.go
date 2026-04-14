package datagsm_test

import (
	"context"
	"net/http"
	"testing"

	datagsm "github.com/jihoonwjj/datagsm-openapi-sdk-go"
)

func TestClubs_List(t *testing.T) {
	want := datagsm.ClubResponse{
		Clubs: []datagsm.ClubDetail{
			{
				Club: datagsm.Club{
					ID: 1, Name: "개발팀", Type: datagsm.ClubTypeMajor, Status: datagsm.ClubStatusActive, FoundedYear: 2020,
				},
				Participants: []datagsm.ParticipantInfo{},
			},
		},
		TotalElements: 1,
		TotalPages:    1,
	}

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/clubs" {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, want)
	}))

	resp, err := client.Clubs().List(context.Background(), datagsm.ClubQuery{Page: 0, Size: 10})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(resp.Clubs) != 1 {
		t.Fatalf("expected 1 club, got %d", len(resp.Clubs))
	}
	if resp.Clubs[0].Name != "개발팀" {
		t.Errorf("expected name '개발팀', got %q", resp.Clubs[0].Name)
	}
}

func TestClubs_Get(t *testing.T) {
	want := datagsm.ClubDetail{
		Club:         datagsm.Club{ID: 5, Name: "AI 연구팀", Type: datagsm.ClubTypeMajor, Status: datagsm.ClubStatusActive, FoundedYear: 2022},
		Participants: []datagsm.ParticipantInfo{},
	}

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/clubs/5" {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, want)
	}))

	got, err := client.Clubs().Get(context.Background(), 5)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got.ID != 5 {
		t.Errorf("expected ID 5, got %d", got.ID)
	}
}
