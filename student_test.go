package datagsm_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	datagsm "github.com/themoment-team/datagsm-openapi-sdk-go"
)

func newTestServer(t *testing.T, handler http.Handler) (*httptest.Server, *datagsm.DataGsmClient) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	c, err := datagsm.NewClient("test-api-key", datagsm.WithBaseURL(srv.URL))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	t.Cleanup(func() { c.Close() })
	return srv, c
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	raw, _ := json.Marshal(data)
	payload := map[string]any{
		"status":  "OK",
		"code":    200,
		"message": "성공",
		"data":    json.RawMessage(raw),
	}
	json.NewEncoder(w).Encode(payload)
}

func TestStudents_List(t *testing.T) {
	want := datagsm.StudentResponse{
		Students: []datagsm.Student{
			{ID: 1, Name: "홍길동", Grade: 1, ClassNum: 1, Number: 1, Major: datagsm.MajorSWDevelopment},
		},
		TotalElements: 1,
		TotalPages:    1,
	}

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/students" {
			http.NotFound(w, r)
			return
		}
		if r.Header.Get("X-API-KEY") != "test-api-key" {
			http.Error(w, `{"message":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
		writeJSON(w, want)
	}))

	resp, err := client.Students().List(context.Background(), datagsm.StudentQuery{Page: 0, Size: 10})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(resp.Students) != 1 {
		t.Fatalf("expected 1 student, got %d", len(resp.Students))
	}
	if resp.Students[0].Name != "홍길동" {
		t.Errorf("expected name '홍길동', got %q", resp.Students[0].Name)
	}
}

func TestStudents_Get(t *testing.T) {
	want := datagsm.Student{ID: 42, Name: "김철수", Grade: 2, Major: datagsm.MajorAI}

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/students/42" {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, want)
	}))

	got, err := client.Students().Get(context.Background(), 42)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got.ID != 42 {
		t.Errorf("expected ID 42, got %d", got.ID)
	}
}

func TestStudents_APIError(t *testing.T) {
	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"message":"invalid api key"}`, http.StatusUnauthorized)
	}))

	_, err := client.Students().List(context.Background(), datagsm.StudentQuery{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var apiErr *datagsm.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *datagsm.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", apiErr.StatusCode)
	}
}

func TestStudents_QueryParams(t *testing.T) {
	var gotParams map[string]string

	_, client := newTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotParams = map[string]string{
			"grade":    r.URL.Query().Get("grade"),
			"classNum": r.URL.Query().Get("classNum"),
			"sortBy":   r.URL.Query().Get("sortBy"),
		}
		writeJSON(w, datagsm.StudentResponse{})
	}))

	_, err := client.Students().List(context.Background(), datagsm.StudentQuery{
		Grade:    datagsm.Ptr(1),
		ClassNum: datagsm.Ptr(2),
		SortBy:   datagsm.StudentSortByName,
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if gotParams["grade"] != "1" {
		t.Errorf("grade: expected '1', got %q", gotParams["grade"])
	}
	if gotParams["classNum"] != "2" {
		t.Errorf("classNum: expected '2', got %q", gotParams["classNum"])
	}
	if gotParams["sortBy"] != "NAME" {
		t.Errorf("sortBy: expected 'NAME', got %q", gotParams["sortBy"])
	}
}
