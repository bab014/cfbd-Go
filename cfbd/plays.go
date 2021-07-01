package cfbd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bab014/go_cfbd/models"
	"github.com/bab014/go_cfbd/opts"
)

func (cfbd *CFBD) GetPlays(queryParams opts.PlaysAPI) ([]models.Plays, error) {

	endpnts := baseUrl.ResolveReference(&url.URL{Path: "plays"})

	req, err := http.NewRequest(http.MethodGet, endpnts.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.PlaysQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var playsData []models.Plays
		if err := json.NewDecoder(res.Body).Decode(&playsData); err != nil {
			return nil, err
		}
		return playsData, nil
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
