package types

type PUUIDCacheCheck struct {
    Found       bool
    PUUID       string
    Stale       bool
}