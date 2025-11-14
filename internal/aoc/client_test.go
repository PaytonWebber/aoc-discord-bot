package aoc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const mockLeaderboardJSON = `{
	"event": "2024",
	"owner_id": 12345,
	"members": {
		"67890": {
			"id": 67890,
			"last_star_ts": 1672444800,
			"global_score": 100,
			"local_score": 200,
			"name": "Test User",
			"completion_day_level": {
				"1": {
					"1": {
						"get_star_ts": 1672444800,
						"star_index": 1
					},
					"2": {
						"get_star_ts": 1672444900,
						"star_index": 2
					}
				}
			},
			"stars": 2
		}
	}
}`

func TestGetLeaderboardSuccess(t *testing.T) {
	// Create a mock server to return a valid JSON response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		expectedPath := "/2024/leaderboard/private/view/test-leaderboard.json"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected URL path %s, got %s", expectedPath, r.URL.Path)
		}
		// Return mock JSON
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, mockLeaderboardJSON)
	}))
	defer mockServer.Close()

	client := NewClient("test-session-cookie", 2024)
	// Inject mock HTTP client
	client.SetHTTPClient(mockServer.Client())

	// Override the request URL using a custom transport
	client.HTTPClient.Transport = rewriteURLTransport("https://adventofcode.com", mockServer.URL)

	// Call the method under test
	leaderboard, err := client.GetLeaderboard("test-leaderboard")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify top-level fields
	if leaderboard.Event != "2024" {
		t.Errorf("Expected event '2024', got '%s'", leaderboard.Event)
	}
	if leaderboard.OwnerID != 12345 {
		t.Errorf("Expected owner ID 12345, got %d", leaderboard.OwnerID)
	}

	// Verify member data
	member, exists := leaderboard.Members["67890"]
	if !exists {
		t.Fatalf("Expected member with ID '67890' in leaderboard members")
	}
	if member.ID != 67890 {
		t.Errorf("Expected member ID 67890, got %d", member.ID)
	}
	if member.Name != "Test User" {
		t.Errorf("Expected member name 'Test User', got '%s'", member.Name)
	}
	if member.Stars != 2 {
		t.Errorf("Expected stars 2, got %d", member.Stars)
	}
	if member.LocalScore != 200 {
		t.Errorf("Expected local score 200, got %d", member.LocalScore)
	}

	// Verify completion day levels
	day1, exists := member.CompletionDayLevels["1"]
	if !exists {
		t.Fatalf("Expected day '1' in completion_day_level")
	}
	if day1.Level1 == nil || day1.Level1.GetStarTs != 1672444800 {
		t.Errorf("Expected Level1 get_star_ts 1672444800, got %v", day1.Level1)
	}
	if day1.Level2 == nil || day1.Level2.GetStarTs != 1672444900 {
		t.Errorf("Expected Level2 get_star_ts 1672444900, got %v", day1.Level2)
	}
}

func TestGetLeaderboardInvalidSession(t *testing.T) {
	// Create a mock server that returns unauthorized response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}))
	defer mockServer.Close()

	client := NewClient("invalid-session-cookie", 2024)
	client.SetHTTPClient(mockServer.Client())

	// Override the request URL
	client.HTTPClient.Transport = rewriteURLTransport("https://adventofcode.com", mockServer.URL)

	_, err := client.GetLeaderboard("test-leaderboard")
	if err == nil {
		t.Fatalf("Expected an error due to unauthorized access, but got none")
	}
	expectedError := "error unmarshalling response body"
	if err.Error()[:len(expectedError)] != expectedError {
		t.Errorf("Expected error starting with '%s', got '%v'", expectedError, err)
	}
}

func TestGetLeaderboardInvalidJSON(t *testing.T) {
	// Create a mock server that returns invalid JSON
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "invalid json")
	}))
	defer mockServer.Close()

	client := NewClient("test-session-cookie", 2024)
	client.SetHTTPClient(mockServer.Client())

	// Override the request URL
	client.HTTPClient.Transport = rewriteURLTransport("https://adventofcode.com", mockServer.URL)

	_, err := client.GetLeaderboard("test-leaderboard")
	if err == nil {
		t.Fatalf("Expected an error due to invalid JSON, but got none")
	}
	expectedError := "error unmarshalling response body"
	if err.Error()[:len(expectedError)] != expectedError {
		t.Errorf("Expected error starting with '%s', got '%v'", expectedError, err)
	}
}

// rewriteURLTransport modifies the request URL to point to the mock server
func rewriteURLTransport(originalBase, mockBase string) http.RoundTripper {
	return &urlRewritingTransport{
		originalBase: originalBase,
		mockBase:     mockBase,
		original:     http.DefaultTransport,
	}
}

type urlRewritingTransport struct {
	originalBase string
	mockBase     string
	original     http.RoundTripper
}

func (t *urlRewritingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Replace the original base URL with the mock server's URL
	if req.URL.String()[:len(t.originalBase)] == t.originalBase {
		req.URL.Scheme = "http"
		req.URL.Host = t.mockBase[len("http://"):]
	}
	return t.original.RoundTrip(req)
}
