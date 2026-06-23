package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAsciiArtHandler(t *testing.T) {
	t.Run("valid POST request", func(t *testing.T) {
		// Create form data
		form := strings.NewReader("text=hello&banner=standard")
		req := httptest.NewRequest("POST", "/ascii-art", form)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		asciiArtHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected 200 OK, got %v", resp.StatusCode)
		}

		// Parse JSON response
		var data map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&data)

		if data["message"] != "received" {
			t.Errorf("Expected message 'received', got %v", data["message"])
		}
		if data["text"] != "hello" {
			t.Errorf("Expected text 'hello', got %v", data["text"])
		}
		if data["banner"] != "standard" {
			t.Errorf("Expected banner 'standard', got %v", data["banner"])
		}
	})

	t.Run("missing text field", func(t *testing.T) {
		form := strings.NewReader("banner=standard")
		req := httptest.NewRequest("POST", "/ascii-art", form)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		asciiArtHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected 400 Bad Request, got %v", resp.StatusCode)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read body: %v", err)
		}
		body := string(bodyBytes)

		if !strings.Contains(body, "text is required") {
			t.Errorf("Expected 'text is required', got %v", body)
		}
	})

	t.Run("missing banner field", func(t *testing.T) {
		form := strings.NewReader("text=hello")
		req := httptest.NewRequest("POST", "/ascii-art", form)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		asciiArtHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected 400 Bad Request, got %v", resp.StatusCode)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read body: %v", err)
		}
		body := string(bodyBytes)

		if !strings.Contains(body, "banner is required") {
			t.Errorf("Expected 'banner is required', got %v", body)
		}
	})

	t.Run("invalid banner", func(t *testing.T) {
		form := strings.NewReader("text=hello&banner=invalid")
		req := httptest.NewRequest("POST", "/ascii-art", form)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		asciiArtHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected 400 Bad Request, got %v", resp.StatusCode)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read body: %v", err)
		}
		body := string(bodyBytes)

		if !strings.Contains(body, "invalid banner") {
			t.Errorf("Expected 'invalid banner', got %v", body)
		}
	})

	t.Run("GET request should fail", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ascii-art", nil)
		w := httptest.NewRecorder()

		asciiArtHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected 405 Method Not Allowed, got %v", resp.StatusCode)
		}
	})
}
