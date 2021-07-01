package opts

import (
	"fmt"
	"net/url"
)

type GamesAPI struct {
	Year       uint
	Week       uint
	SeasonType string
	Team       string
	Home       string
	Away       string
	Conference string
	Id         uint
	GameId     uint
	Category   string
}

func (gameArgs GamesAPI) GamesQuery() url.Values {
	q := make(url.Values)

	if gameArgs.Year != 0 {
		q.Add("year", fmt.Sprint(gameArgs.Year))
	}
	if gameArgs.Week != 0 {
		q.Add("week", fmt.Sprint(gameArgs.Week))
	}
	if gameArgs.SeasonType == "" {
		q.Add("seasonType", "regular")
	} else {
		q.Add("seasonType", gameArgs.SeasonType)
	}
	if gameArgs.Team != "" {
		q.Add("team", gameArgs.Team)
	}
	if gameArgs.Home != "" {
		q.Add("home", gameArgs.Home)
	}
	if gameArgs.Away != "" {
		q.Add("away", gameArgs.Away)
	}
	if gameArgs.Conference != "" {
		q.Add("conference", gameArgs.Conference)
	}
	if gameArgs.Id != 0 {
		q.Add("id", fmt.Sprint(gameArgs.Id))
	}
	if gameArgs.GameId != 0 {
		q.Add("gameId", fmt.Sprint(gameArgs.GameId))
	}
	if gameArgs.Category != "" {
		q.Add("category", gameArgs.Category)
	}

	return q
}
