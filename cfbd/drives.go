package cfbd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bab014/go_cfbd/models"
)

func (cfbd *CFBD) GetDrives(queryParams DrivesAPI) ([]models.Drives, error) {

	endpnt := baseUrl.ResolveReference(&url.URL{Path: "drives"})
	req, err := http.NewRequest(http.MethodGet, endpnt.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", cfbd.token))
	req.URL.RawQuery = queryParams.DrivesQuery().Encode()

	res, err := cfbd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var drivesData []models.Drives
		if err := json.NewDecoder(res.Body).Decode(&drivesData); err != nil {
			return nil, err
		}
		return drivesData, nil
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
