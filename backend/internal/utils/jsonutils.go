package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"winnable/internal/types"
)

func WriteProfileToFile(data types.LeagueProfilePage, riotID string) error {
	path := fmt.Sprintf("./devfiles/profiles/%s_LoLProfile.JSON", riotID)
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create profile file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("failed to write profile JSON: %w", err)
	}

	log.Printf("wrote %s profile to %s", riotID, path)
	return nil
}

func WriteMasteryToFile(data types.MasteryData, riotID string) error {
	path := fmt.Sprintf("./devfiles/masteries/%s_MasteryData.JSON", riotID)
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create mastery file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("failed to write mastery JSON: %w", err)
	}

	log.Printf("wrote %s mastery to %s", riotID, path)
	return nil
}
