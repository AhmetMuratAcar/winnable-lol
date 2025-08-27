package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"winnable/internal/types"
)

func WriteProfileToFile(data types.LeagueProfilePage, riotID string) error {
	path := fmt.Sprintf("./devfiles/%s_LoLProfile.JSON", riotID)
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	log.Printf("wrote %s profile to %s", riotID, path)
	return nil
}
