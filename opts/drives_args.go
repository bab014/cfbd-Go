package opts

import (
	"fmt"
	"net/url"
)

type DrivesAPI struct {
	Year              uint
	Week              uint
	SeasonType        string
	Team              string
	Offense           string
	Defense           string
	Conference        string
	OffenseConference string
	DefenseConference string
}

func (drivesArgs DrivesAPI) DrivesQuery() url.Values {
	q := make(url.Values)

	if drivesArgs.Year != 0 {
		q.Add("year", fmt.Sprint(drivesArgs.Year))
	}
	if drivesArgs.Week != 0 {
		q.Add("week", fmt.Sprint(drivesArgs.Week))
	}
	if drivesArgs.SeasonType == "" {
		q.Add("seasonType", "regular")
	} else {
		q.Add("seasonType", drivesArgs.SeasonType)
	}
	if drivesArgs.Team != "" {
		q.Add("team", drivesArgs.Team)
	}
	if drivesArgs.Offense != "" {
		q.Add("offense", drivesArgs.Offense)
	}
	if drivesArgs.Defense != "" {
		q.Add("defense", drivesArgs.Defense)
	}
	if drivesArgs.OffenseConference != "" {
		q.Add("offenseConference", drivesArgs.OffenseConference)
	}
	if drivesArgs.DefenseConference != "" {
		q.Add("defenseConference", drivesArgs.DefenseConference)
	}
	if drivesArgs.Conference != "" {
		q.Add("conference", drivesArgs.Conference)
	}

	return q
}
