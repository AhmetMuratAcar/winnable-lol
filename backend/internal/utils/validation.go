package utils

import (
	"encoding/json"
	"net/http"
	"winnable/internal/types"
)

type ErrorResponse struct {
	Error         string   `json:"error"`
	MissingParams []string `json:"missingParams,omitempty"`
}

func ValidateLeagueProfileReq(w http.ResponseWriter, r *http.Request) (types.RequestBody, bool) {
	req := types.RequestBody{
		GameName: r.URL.Query().Get("gameName"),
		TagLine:  r.URL.Query().Get("tagLine"),
		Region:   r.URL.Query().Get("region"),
	}

	missingParams := []string{}
	if req.GameName == "" {
		missingParams = append(missingParams, "gameName")
	}
	if req.TagLine == "" {
		missingParams = append(missingParams, "tagLine")
	}
	if req.Region == "" {
		missingParams = append(missingParams, "region")
	}

	if len(missingParams) > 0 {
		resp := ErrorResponse{
			Error:         "missing required query parameters",
			MissingParams: missingParams,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(resp)
		return types.RequestBody{}, false
	}

	return req, true
}
