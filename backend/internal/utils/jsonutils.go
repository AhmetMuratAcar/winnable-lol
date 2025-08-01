package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"winnable/internal/types"
)

func WriteMasteryToFile(data []types.ChampionMastery, riotID string) error {
	path := fmt.Sprintf("./devfiles/%sChampionMasteries.JSON", riotID)
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

	log.Printf("wrote %d entries to %s", len(data), path)
	return nil
}
