package opts

import (
	"fmt"
	"net/url"
)

type TeamsAPI struct {
	Year    uint
	School  *string
	Team1   string
	Team2   string
	MinYear uint
	MaxYear uint
}

func (teamArgs TeamsAPI) TeamsQuery() url.Values {
	q := make(url.Values)

	if teamArgs.Year != 0 {
		q.Add("year", fmt.Sprint(teamArgs.Year))
	}
	if teamArgs.School != nil {
		q.Add("school", *teamArgs.School)
	}
	if teamArgs.Team1 != "" {
		q.Add("team1", teamArgs.Team1)
	}
	if teamArgs.Team2 != "" {
		q.Add("team2", teamArgs.Team2)
	}
	if teamArgs.MinYear != 0 {
		q.Add("minYear", fmt.Sprint(teamArgs.MinYear))
	}
	if teamArgs.MaxYear != 0 {
		q.Add("maxYear", fmt.Sprint(teamArgs.MaxYear))
	}

	return q
}
