package cfbd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bab014/go_cfbd/models"
	"github.com/bab014/go_cfbd/opts"
)

func (cfbd *CFBD) GetTeams(queryParams opts.TeamsAPI) ([]models.Team, error) {

	endpt := baseUrl.ResolveReference(&url.URL{Path: "teams/fbs"})

	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.TeamsQuery().Encode()

	// Authenticate the request
	// add queryParams to the request

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var teamsData []models.Team

		// Deserialize success or error response and return its data
		if err := json.NewDecoder(res.Body).Decode(&teamsData); err != nil {
			return nil, err
		}
		return teamsData, nil
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

func (cfbd *CFBD) GetRoster(queryParams opts.TeamsAPI) ([]models.Players, error) {

	endpnt := baseUrl.ResolveReference(&url.URL{Path: "roster"})

	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.TeamsQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var rosterData []models.Players

		// Deserialize success or error response and return its data
		if err := json.NewDecoder(res.Body).Decode(&rosterData); err != nil {
			return nil, err
		}
		return rosterData, nil
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

func (cfbd *CFBD) GetTalent(queryParams opts.TeamsAPI) ([]models.Talent, error) {

	endpnt := baseUrl.ResolveReference(&url.URL{Path: "talent"})

	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.TeamsQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var talentData []models.Talent

		// Deserialize success or error response and return its data
		if err := json.NewDecoder(res.Body).Decode(&talentData); err != nil {
			return nil, err
		}
		return talentData, nil
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

func (cfbd *CFBD) GetMatchup(queryParams opts.TeamsAPI) (*models.Matchup, error) {

	endpnt := baseUrl.ResolveReference(&url.URL{Path: "teams/matchup"})

	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.TeamsQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var matchupData models.Matchup

		// Deserialize success or error response and return its data
		if err := json.NewDecoder(res.Body).Decode(&matchupData); err != nil {
			return nil, err
		}
		return &matchupData, nil
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
