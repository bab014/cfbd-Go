package models

import "time"

type Games struct {
	ID              int         `json:"id"`
	Season          int         `json:"season"`
	Week            int         `json:"week"`
	SeasonType      string      `json:"season_type"`
	StartDate       time.Time   `json:"start_date"`
	StartTimeTbd    interface{} `json:"start_time_tbd"`
	NeutralSite     bool        `json:"neutral_site"`
	ConferenceGame  bool        `json:"conference_game"`
	Attendance      int         `json:"attendance"`
	VenueID         int         `json:"venue_id"`
	Venue           string      `json:"venue"`
	HomeID          int         `json:"home_id"`
	HomeTeam        string      `json:"home_team"`
	HomeConference  string      `json:"home_conference"`
	HomePoints      int         `json:"home_points"`
	HomeLineScores  []int       `json:"home_line_scores"`
	HomePostWinProb string      `json:"home_post_win_prob"`
	AwayID          int         `json:"away_id"`
	AwayTeam        string      `json:"away_team"`
	AwayConference  string      `json:"away_conference"`
	AwayPoints      int         `json:"away_points"`
	AwayLineScores  []int       `json:"away_line_scores"`
	AwayPostWinProb string      `json:"away_post_win_prob"`
	ExcitementIndex string      `json:"excitement_index"`
	Highlights      string      `json:"highlights"`
}

type Records struct {
	Year       int    `json:"year"`
	Team       string `json:"team"`
	Conference string `json:"conference"`
	Division   string `json:"division"`
	Total      struct {
		Games  int `json:"games"`
		Wins   int `json:"wins"`
		Losses int `json:"losses"`
		Ties   int `json:"ties"`
	} `json:"total"`
	ConferenceGames struct {
		Games  int `json:"games"`
		Wins   int `json:"wins"`
		Losses int `json:"losses"`
		Ties   int `json:"ties"`
	} `json:"conferenceGames"`
	HomeGames struct {
		Games  int `json:"games"`
		Wins   int `json:"wins"`
		Losses int `json:"losses"`
		Ties   int `json:"ties"`
	} `json:"homeGames"`
	AwayGames struct {
		Games  int `json:"games"`
		Wins   int `json:"wins"`
		Losses int `json:"losses"`
		Ties   int `json:"ties"`
	} `json:"awayGames"`
}

type Calendar struct {
	Season         string `json:"season"`
	Week           int    `json:"week"`
	SeasonType     string `json:"seasonType"`
	FirstGameStart string `json:"firstGameStart"`
	LastGameStart  string `json:"lastGameStart"`
}

type Media struct {
	ID             int    `json:"id"`
	Season         int    `json:"season"`
	Week           int    `json:"week"`
	SeasonType     string `json:"seasonType"`
	StartTime      string `json:"startTime"`
	IsStartTimeTBD bool   `json:"isStartTimeTBD"`
	HomeTeam       string `json:"homeTeam"`
	HomeConference string `json:"homeConference"`
	AwayTeam       string `json:"awayTeam"`
	AwayConference string `json:"awayConference"`
	MediaType      string `json:"mediaType"`
	Outlet         string `json:"outlet"`
}

type GamesPlayers struct {
	ID    int `json:"id"`
	Teams []struct {
		School     string `json:"school"`
		Conference string `json:"conference"`
		HomeAway   string `json:"homeAway"`
		Points     int    `json:"points"`
		Categories []struct {
			Name  string `json:"name"`
			Types []struct {
				Name     string `json:"name"`
				Athletes []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Stat string `json:"stat"`
				} `json:"athletes"`
			} `json:"types"`
		} `json:"categories"`
	} `json:"teams"`
}

type GamesTeams struct {
	ID    int `json:"id"`
	Teams []struct {
		School     string `json:"school"`
		Conference string `json:"conference"`
		HomeAway   string `json:"homeAway"`
		Points     int    `json:"points"`
		Stats      []struct {
			Category string `json:"category"`
			Stat     string `json:"stat"`
		} `json:"stats"`
	} `json:"teams"`
}

type GamesBoxAdvanced struct {
	Teams struct {
		Ppa []struct {
			Team    string `json:"team"`
			Plays   int    `json:"plays"`
			Overall struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"overall"`
			Passing struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"passing"`
			Rushing struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"rushing"`
		} `json:"ppa"`
		CumulativePpa []struct {
			Team    string `json:"team"`
			Plays   int    `json:"plays"`
			Overall struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"overall"`
			Passing struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"passing"`
			Rushing struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"rushing"`
		} `json:"cumulativePpa"`
		SuccessRates []struct {
			Team    string `json:"team"`
			Overall struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"overall"`
			StandardDowns struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"standardDowns"`
			PassingDowns struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"passingDowns"`
		} `json:"successRates"`
		Explosiveness []struct {
			Team    string `json:"team"`
			Overall struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
			} `json:"overall"`
		} `json:"explosiveness"`
		Rushing []struct {
			Team                    string `json:"team"`
			PowerSuccess            string `json:"powerSuccess"`
			StuffRate               string `json:"stuffRate"`
			LineYards               string `json:"lineYards"`
			LineYardsAverage        string `json:"lineYardsAverage"`
			SecondLevelYards        string `json:"secondLevelYards"`
			SecondLevelYardsAverage string `json:"secondLevelYardsAverage"`
			OpenFieldYards          string `json:"openFieldYards"`
			OpenFieldYardsAverage   string `json:"openFieldYardsAverage"`
		} `json:"rushing"`
		Havoc []struct {
			Team       string `json:"team"`
			Total      string `json:"total"`
			FrontSeven string `json:"frontSeven"`
			Db         string `json:"db"`
		} `json:"havoc"`
		ScoringOpportunities []struct {
			Team                 string  `json:"team"`
			Opportunities        int     `json:"opportunities"`
			Points               int     `json:"points"`
			PointsPerOpportunity float64 `json:"pointsPerOpportunity"`
		} `json:"scoringOpportunities"`
		FieldPosition []struct {
			Team                           string `json:"team"`
			AverageStart                   string `json:"averageStart"`
			AverageStartingPredictedPoints string `json:"averageStartingPredictedPoints"`
		} `json:"fieldPosition"`
	} `json:"teams"`
	Players struct {
		Usage []struct {
			Player   string  `json:"player"`
			Team     string  `json:"team"`
			Position string  `json:"position"`
			Total    float64 `json:"total"`
			Quarter1 float64 `json:"quarter1"`
			Quarter2 float64 `json:"quarter2"`
			Quarter3 float64 `json:"quarter3"`
			Quarter4 float64 `json:"quarter4"`
			Rushing  float64 `json:"rushing"`
			Passing  float64 `json:"passing"`
		} `json:"usage"`
		Ppa []struct {
			Player   string `json:"player"`
			Team     string `json:"team"`
			Position string `json:"position"`
			Average  struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
				Rushing  float64 `json:"rushing"`
				Passing  float64 `json:"passing"`
			} `json:"average"`
			Cumulative struct {
				Total    float64 `json:"total"`
				Quarter1 float64 `json:"quarter1"`
				Quarter2 float64 `json:"quarter2"`
				Quarter3 float64 `json:"quarter3"`
				Quarter4 float64 `json:"quarter4"`
				Rushing  float64 `json:"rushing"`
				Passing  float64 `json:"passing"`
			} `json:"cumulative"`
		} `json:"ppa"`
	} `json:"players"`
}
