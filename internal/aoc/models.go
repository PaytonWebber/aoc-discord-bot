package aoc

// Leaderboard represents the entire JSON structure
type Leaderboard struct {
	Members map[string]Member `json:"members"`
	Event   string            `json:"event"`
	OwnerID int               `json:"owner_id"`
}

// Member represents each member in the leaderboard
type Member struct {
	ID                 int                        `json:"id"`
	LastStarTs         int64                      `json:"last_star_ts"`
	GlobalScore        int                        `json:"global_score"`
	LocalScore         int                        `json:"local_score"`
	Name               string                     `json:"name"`
	CompletionDayLevel map[string]map[string]Star `json:"completion_day_level"`
	Stars              int                        `json:"stars"`
}

// Star represents the completion data for each star
type Star struct {
	GetStarTs int64 `json:"get_star_ts"`
	StarIndex int   `json:"star_index"`
}
