package opts

import (
	"fmt"
	"net/url"
)

type PlaysAPI struct {
	SeasonType        string
	Year              uint
	Week              uint
	Team              string
	Offense           string
	Defense           string
	Conference        string
	OffenseConference string
	DefenseConference string
	PlayType          string
}

func (playsArgs PlaysAPI) PlaysQuery() url.Values {
	q := make(url.Values)

	if playsArgs.Year != 0 {
		q.Add("year", fmt.Sprint(playsArgs.Year))
	}
	if playsArgs.Week != 0 {
		q.Add("week", fmt.Sprint(playsArgs.Week))
	}
	if playsArgs.SeasonType == "" {
		q.Add("seasonType", "regular")
	} else {
		q.Add("seasonType", playsArgs.SeasonType)
	}
	if playsArgs.Team != "" {
		q.Add("team", playsArgs.Team)
	}
	if playsArgs.Offense != "" {
		q.Add("offense", playsArgs.Offense)
	}
	if playsArgs.Defense != "" {
		q.Add("defense", playsArgs.Defense)
	}
	if playsArgs.OffenseConference != "" {
		q.Add("offenseConference", playsArgs.OffenseConference)
	}
	if playsArgs.DefenseConference != "" {
		q.Add("defenseConference", playsArgs.DefenseConference)
	}
	if playsArgs.Conference != "" {
		q.Add("conference", playsArgs.Conference)
	}
	if playsArgs.PlayType != "" {
		q.Add("conference", playsArgs.Conference)
	}

	return q
}
