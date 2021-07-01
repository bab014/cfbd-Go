package models

type Plays struct {
	ID                string `json:"id"`
	DriveID           string `json:"drive_id"`
	GameID            int    `json:"game_id"`
	DriveNumber       int    `json:"drive_number"`
	PlayNumber        int    `json:"play_number"`
	Offense           string `json:"offense"`
	OffenseConference string `json:"offense_conference"`
	OffenseScore      int    `json:"offense_score"`
	Defense           string `json:"defense"`
	Home              string `json:"home"`
	Away              string `json:"away"`
	DefenseConference string `json:"defense_conference"`
	DefenseScore      int    `json:"defense_score"`
	Period            int    `json:"period"`
	Clock             struct {
		Minutes int `json:"minutes"`
		Seconds int `json:"seconds"`
	} `json:"clock"`
	OffenseTimeouts int    `json:"offense_timeouts"`
	DefenseTimeouts int    `json:"defense_timeouts"`
	YardLine        int    `json:"yard_line"`
	YardsToGoal     int    `json:"yards_to_goal"`
	Down            int    `json:"down"`
	Distance        int    `json:"distance"`
	YardsGained     int    `json:"yards_gained"`
	Scoring         bool   `json:"scoring"`
	PlayType        string `json:"play_type"`
	PlayText        string `json:"play_text"`
	Ppa             string `json:"ppa"`
}
