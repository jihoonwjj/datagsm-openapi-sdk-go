package datagsm_test

import (
	"testing"
	"time"

	datagsm "github.com/themoment-team/datagsm-openapi-sdk-go"
)

func TestNewClient_EmptyAPIKey(t *testing.T) {
	_, err := datagsm.NewClient("")
	if err == nil {
		t.Fatal("expected error for empty apiKey, got nil")
	}
}

func TestNewClient_ValidAPIKey(t *testing.T) {
	c, err := datagsm.NewClient("test-key")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer c.Close()

	if c.Students() == nil {
		t.Error("Students() returned nil")
	}
	if c.Clubs() == nil {
		t.Error("Clubs() returned nil")
	}
	if c.Projects() == nil {
		t.Error("Projects() returned nil")
	}
	if c.NEIS() == nil {
		t.Error("NEIS() returned nil")
	}
}

func TestNewClient_WithOptions(t *testing.T) {
	c, err := datagsm.NewClient("test-key",
		datagsm.WithBaseURL("https://custom.example.com"),
		datagsm.WithTimeout(10*time.Second),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer c.Close()
}
