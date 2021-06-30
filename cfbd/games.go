package cfbd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/bab014/go_cfbd/models"
)

func (cfbd *CFBD) GetGames(queryParams GamesAPI) ([]models.Games, error) {

	endpt := baseUrl.ResolveReference(&url.URL{Path: "games"})

	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.GamesQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var gamesData []models.Games
		if err := json.NewDecoder(res.Body).Decode(&gamesData); err != nil {
			return nil, err
		}
		return gamesData, nil
	case 400:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// Returns Records data
// Query options are:
// Year       uint;
// Week       uint;
// SeasonType string;
// Team       string;
// Home       string;
// Away       string;
// Conference string;
// Id         uint
func (cfbd *CFBD) GetRecords(queryParams GamesAPI) ([]models.Records, error) {
	endpnt := baseUrl.ResolveReference(&url.URL{Path: "records"})

	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.GamesQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var recordsData []models.Records
		if err := json.NewDecoder(res.Body).Decode(&recordsData); err != nil {
			return nil, err
		}
		return recordsData, nil
	case 400:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

func (cfbd *CFBD) GetCalendar(queryParams GamesAPI) ([]models.Calendar, error) {
	endpnt := baseUrl.ResolveReference(&url.URL{Path: "calendar"})

	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.GamesQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var calendarData []models.Calendar
		if err := json.NewDecoder(res.Body).Decode(&calendarData); err != nil {
			return nil, err
		}
		return calendarData, nil

	case 400:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

func (cfbd *CFBD) GetMedia(queryParams GamesAPI) ([]models.Media, error) {

	endpnt := baseUrl.ResolveReference(&url.URL{Path: "games/media"})
	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.GamesQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var mediaData []models.Media
		if err := json.NewDecoder(res.Body).Decode(&mediaData); err != nil {
			return nil, err
		}
		return mediaData, nil
	case 400:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

func (cfbd *CFBD) GetGamesPlayers(queryParams GamesAPI) ([]models.GamesPlayers, error) {

	endpnt := baseUrl.ResolveReference(&url.URL{Path: "games/players"})
	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.GamesQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var gamesPlayersData []models.GamesPlayers
		if err := json.NewDecoder(res.Body).Decode(&gamesPlayersData); err != nil {
			return nil, err
		}
		return gamesPlayersData, nil
	case 400:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

func (cfbd *CFBD) GetGamesTeams(queryParams GamesAPI) ([]models.GamesTeams, error) {

	endpnt := baseUrl.ResolveReference(&url.URL{Path: "games/teams"})
	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.GamesQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var gamesTeamsData []models.GamesTeams
		if err := json.NewDecoder(res.Body).Decode(&gamesTeamsData); err != nil {
			return nil, err
		}
		return gamesTeamsData, nil
	case 400:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

func (cfbd *CFBD) GetGamesBoxAdvanced(queryParams GamesAPI) (*models.GamesBoxAdvanced, error) {

	endpnt := baseUrl.ResolveReference(&url.URL{Path: "game/box/advanced"})
	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.GamesQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var gamesBoxAdvancedData models.GamesBoxAdvanced
		if err := json.NewDecoder(res.Body).Decode(&gamesBoxAdvancedData); err != nil {
			return nil, err
		}
		return &gamesBoxAdvancedData, nil
	case 400:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}
