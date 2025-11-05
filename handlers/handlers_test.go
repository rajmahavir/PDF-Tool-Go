package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleHome)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HandleHome returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("HandleHome returned wrong content type: got %v want %v", contentType, "text/html")
	}

	body := rr.Body.String()
	expectedStrings := []string{
		"PDF Tools",
		"Remove Pages",
		"Merge PDFs",
	}

	for _, expected := range expectedStrings {
		if !contains(body, expected) {
			t.Errorf("HandleHome body does not contain expected string: %s", expected)
		}
	}
}

func TestHandleCredits(t *testing.T) {
	req, err := http.NewRequest("GET", "/credits", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleCredits)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HandleCredits returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("HandleCredits returned wrong content type: got %v want %v", contentType, "text/html")
	}

	body := rr.Body.String()
	expectedStrings := []string{
		"Credits",
		"Claude.ai",
		"pdfcpu",
		"MIT License",
	}

	for _, expected := range expectedStrings {
		if !contains(body, expected) {
			t.Errorf("HandleCredits body does not contain expected string: %s", expected)
		}
	}
}

func TestHandleMergePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/merge", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleMergePage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HandleMergePage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := rr.Body.String()
	expectedStrings := []string{
		"Merge PDFs",
		"First PDF",
		"Second PDF",
	}

	for _, expected := range expectedStrings {
		if !contains(body, expected) {
			t.Errorf("HandleMergePage body does not contain expected string: %s", expected)
		}
	}
}

func TestHandleRemovePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/remove", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRemovePage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HandleRemovePage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := rr.Body.String()
	expectedStrings := []string{
		"Remove Pages",
		"Upload PDF",
	}

	for _, expected := range expectedStrings {
		if !contains(body, expected) {
			t.Errorf("HandleRemovePage body does not contain expected string: %s", expected)
		}
	}
}

func TestHandlePDFInfoInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/pdfinfo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePDFInfo)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("HandlePDFInfo with GET should return %v, got %v", http.StatusMethodNotAllowed, status)
	}
}

func TestHandleMergeInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/merge-pdfs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleMerge)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("HandleMerge with GET should return %v, got %v", http.StatusMethodNotAllowed, status)
	}
}

func TestHandleRemovePagesInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/remove-pages", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRemovePages)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("HandleRemovePages with GET should return %v, got %v", http.StatusMethodNotAllowed, status)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
