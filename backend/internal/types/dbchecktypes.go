package types

import "time"

type PUUIDCacheCheck struct {
    Found       bool
    PUUID       string
    Stale       bool
    LastUpdated time.Time
}