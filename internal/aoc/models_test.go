package aoc

import (
	"encoding/json"
	"testing"
)

func TestLeaderboardUnmarshal(t *testing.T) {
	// Sample JSON response
	mockJSON := `{
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

	// Attempt to unmarshal the JSON into the Leaderboard struct
	var leaderboard Leaderboard
	err := json.Unmarshal([]byte(mockJSON), &leaderboard)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the top-level fields
	if leaderboard.Event != "2024" {
		t.Errorf("Expected Event '2024', got '%s'", leaderboard.Event)
	}
	if leaderboard.OwnerID != 12345 {
		t.Errorf("Expected OwnerID 12345, got %d", leaderboard.OwnerID)
	}

	// Verify the members map
	member, exists := leaderboard.Members["67890"]
	if !exists {
		t.Fatalf("Member '67890' not found in leaderboard members")
	}

	if member.ID != 67890 {
		t.Errorf("Expected Member ID 67890, got %d", member.ID)
	}
	if member.Name != "Test User" {
		t.Errorf("Expected Member Name 'Test User', got '%s'", member.Name)
	}
	if member.Stars != 2 {
		t.Errorf("Expected Member Stars 2, got %d", member.Stars)
	}
	if member.LastStarTs != 1672444800 {
		t.Errorf("Expected LastStarTs 1672444800, got %d", member.LastStarTs)
	}

	// Verify nested CompletionDayLevels
	completion, exists := member.CompletionDayLevels["1"]
	if !exists {
		t.Fatalf("Day '1' not found in completion_day_level")
	}

	if completion.Level1 == nil || completion.Level1.GetStarTs != 1672444800 {
		t.Errorf("Expected Level1 GetStarTs 1672444800, got %v", completion.Level1)
	}
	if completion.Level2 == nil || completion.Level2.GetStarTs != 1672444900 {
		t.Errorf("Expected Level2 GetStarTs 1672444900, got %v", completion.Level2)
	}
}

func TestMemberUnmarshalWithMissingFields(t *testing.T) {
	// JSON with missing optional fields
	mockJSON := `{
		"id": 67890,
		"last_star_ts": 1672444800,
		"global_score": 100,
		"local_score": 200,
		"name": "Test User",
		"stars": 2
	}`

	var member Member
	err := json.Unmarshal([]byte(mockJSON), &member)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the parsed member
	if member.ID != 67890 {
		t.Errorf("Expected Member ID 67890, got %d", member.ID)
	}
	if member.Name != "Test User" {
		t.Errorf("Expected Member Name 'Test User', got '%s'", member.Name)
	}
	if member.Stars != 2 {
		t.Errorf("Expected Member Stars 2, got %d", member.Stars)
	}

	// Check that missing fields have default values
	if member.CompletionDayLevels != nil {
		t.Errorf("Expected CompletionDayLevels to be nil, got %v", member.CompletionDayLevels)
	}
}
