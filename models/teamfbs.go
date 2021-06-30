package models

type Location struct {
	VenueId         int     `json:"vanue_id"`
	Name            string  `json:"name"`
	City            string  `json:"city"`
	State           string  `json:"state"`
	Zip             string  `json:"zip"`
	CountryCode     string  `json:"country_code"`
	Timezone        string  `json:"timezone"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	Elevation       string  `json:"elevation"`
	Capacity        int     `json:"capacity"`
	YearConstructed int     `json:"year_constructed"`
	Grass           bool    `json:"grass"`
	Dome            bool    `json:"dome"`
}

type Team struct {
	Id           int      `json:"id"`
	School       string   `json:"school"`
	Mascot       string   `json:"mascot"`
	Abbreviation string   `json:"abbreviation"`
	AltName1     string   `json:"alt_name_1"`
	AltName2     string   `json:"alt_name_2"`
	AltName3     string   `json:"alt_name_3"`
	Conference   string   `json:"conference"`
	Division     string   `json:"division"`
	Color        string   `json:"color"`
	AltColor     string   `json:"alt_color"`
	Logos        []string `json:"logos"`
	Location     Location `json:"location"`
}

type Players struct {
	ID             string `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Team           string `json:"team"`
	Weight         int    `json:"weight"`
	Height         int    `json:"height"`
	Jersey         int    `json:"jersey"`
	Year           int    `json:"year"`
	Position       string `json:"position"`
	HomeCity       string `json:"home_city"`
	HomeState      string `json:"home_state"`
	HomeCountry    string `json:"home_country"`
	HomeLatitude   string `json:"home_latitude"`
	HomeLongitude  string `json:"home_longitude"`
	HomeCountyFips string `json:"home_county_fips"`
	RecruitIds     []int  `json:"recruit_ids"`
}

type Talent struct {
	Year   int    `json:"year"`
	School string `json:"school"`
	Talent string `json:"talent"`
}

type Matchup struct {
	Team1     string `json:"team1"`
	Team2     string `json:"team2"`
	StartYear string `json:"startYear"`
	EndYear   string `json:"endYear"`
	Team1Wins int    `json:"team1Wins"`
	Team2Wins int    `json:"team2Wins"`
	Ties      int    `json:"ties"`
	Games     []struct {
		Season      int    `json:"season"`
		Week        int    `json:"week"`
		SeasonType  string `json:"season_type"`
		Date        string `json:"date"`
		NeutralSite bool   `json:"neutralSite"`
		Venue       string `json:"venue"`
		HomeTeam    string `json:"homeTeam"`
		HomeScore   int    `json:"homeScore"`
		AwayTeam    string `json:"awayTeam"`
		AwayScore   int    `json:"awayScore"`
		Winner      string `json:"winner"`
	} `json:"games"`
}
