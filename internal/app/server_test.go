package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHomePageRendersMachineList(t *testing.T) {
	server, err := NewServer()
	if err != nil {
		t.Fatalf("NewServer returned error: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	server.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	body := rec.Body.String()
	if !strings.Contains(body, "洗濯機一覧") {
		t.Fatalf("expected machine list heading in response: %s", body)
	}
}

func TestStartSessionFragmentReturnsSuccessMessage(t *testing.T) {
	server, err := NewServer()
	if err != nil {
		t.Fatalf("NewServer returned error: %v", err)
	}

	form := url.Values{}
	form.Set("machine_id", "washer-a")

	req := httptest.NewRequest(http.MethodPost, "/sessions", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()
	server.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	body := rec.Body.String()
	if !strings.Contains(body, "利用開始を受け付けました") {
		t.Fatalf("expected success message in response: %s", body)
	}
}

func TestCurrentSessionPageRendersCurrentSessionHeading(t *testing.T) {
	server, err := NewServer()
	if err != nil {
		t.Fatalf("NewServer returned error: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/sessions/current", nil)
	rec := httptest.NewRecorder()

	server.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	body := rec.Body.String()
	if !strings.Contains(body, "現在のセッション") {
		t.Fatalf("expected current session heading in response: %s", body)
	}

	if !strings.Contains(body, "洗濯物を回収しました") {
		t.Fatalf("expected claim action in response: %s", body)
	}
}
