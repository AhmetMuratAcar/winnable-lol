package types

type PUUIDCacheCheck struct {
	Found bool
	PUUID string
	Stale bool
}

type CachedProfileCheckList struct {
	PUUID       bool
	ProfileIcon bool
	Level       bool
	Ranks       bool
	Masteries   bool
	Matches     bool
}
