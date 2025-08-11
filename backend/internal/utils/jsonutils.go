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
		log.Fatalf("failed to create file: %v", err)
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		log.Fatalf("failed to write JSON: %v", err)
		return err
	}

	log.Printf("wrote %s profile to %s", riotID, path)
	return nil
}
