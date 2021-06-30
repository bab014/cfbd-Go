package cfbd

import "fmt"

type ErrorResponse struct {
	Description string `json:"description"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("API error: %s", err.Description)
}
