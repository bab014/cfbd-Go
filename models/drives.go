package models

type Drives struct {
	Offense           string `json:"offense"`
	OffenseConference string `json:"offense_conference"`
	Defense           string `json:"defense"`
	DefenseConference string `json:"defense_conference"`
	GameID            int    `json:"game_id"`
	ID                string `json:"id"`
	DriveNumber       int    `json:"drive_number"`
	Scoring           bool   `json:"scoring"`
	StartPeriod       int    `json:"start_period"`
	StartYardline     int    `json:"start_yardline"`
	StartYardsToGoal  int    `json:"start_yards_to_goal"`
	StartTime         struct {
		Minutes int `json:"minutes"`
		Seconds int `json:"seconds"`
	} `json:"start_time"`
	EndPeriod      int `json:"end_period"`
	EndYardline    int `json:"end_yardline"`
	EndYardsToGoal int `json:"end_yards_to_goal"`
	EndTime        struct {
		Minutes int `json:"minutes"`
		Seconds int `json:"seconds"`
	} `json:"end_time"`
	Plays             int    `json:"plays"`
	Yards             int    `json:"yards"`
	DriveResult       string `json:"drive_result"`
	IsHomeOffense     bool   `json:"is_home_offense"`
	StartOffenseScore int    `json:"start_offense_score"`
	StartDefenseScore int    `json:"start_defense_score"`
	EndOffenseScore   int    `json:"end_offense_score"`
	EndDefenseScore   int    `json:"end_defense_score"`
}
