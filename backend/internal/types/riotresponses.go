package types

type AccountResponse struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type ChampionMastery struct {
	ChampionID     int `json:"championId"`
	ChampionLevel  int `json:"championLevel"`
	ChampionPoints int `json:"championPoints"`
}

type MasteryData struct {
	ChampionMasteries  []ChampionMastery `json:"championMasteries"`
	TotalMastery       int               `json:"totalMastery"`
	TotalMasteryPoints int               `json:"totalMasteryPoints"`
	ChampionsPlayed    int               `json:"championsPlayed"`
}

type LeagueMatch struct {
	EndOfGameResult    string                   `json:"endOfGameResult"`
	GameDuration       int                      `json:"gameDuration"`
	GameStartTimestamp int                      `json:"gameStartTimeStamp"`
	GameVersion        string                   `json:"gameVersion"`
	MatchID            string                   `json:"matchId"`
	ParticipantPUUIDs  []string                 `json:"participantPUUIDs"`
	Participants       []LeagueMatchParticipant `json:"participants"`
	QueueId            int                      `json:"queueId"`
	Bans               []int                    `json:"bans"`
}

type LeagueMatchParticipant struct {
	Assists                     int    `json:"assists"`
	ChampionID                  int    `json:"championId"`
	ChampLevel                  int    `json:"champLevel"`
	Deaths                      int    `json:"deaths"`
	GoldEarned                  int    `json:"goldEarned"`
	Items                       []int  `json:"items"`
	Kills                       int    `json:"kills"`
	ProfileIconID               int    `json:"profileIcon"`
	PUUID                       string `json:"puuid"`
	RiotIDGameName              string `json:"riotIdGameName"`
	RiotIDTagline               string `json:"riotIdTagline"`
	Summoner1ID                 int    `json:"summoner1Id"`
	Summoner2ID                 int    `json:"summoner2Id"`
	SummonerLevel               int    `json:"summonerLevel"`
	TeamPosition                string `json:"teamPosition"`
	TotalDamageDealtToChampions int    `json:"totalDamageDealtToChampions"`
	TotalMinionsKilled          int    `json:"totalMinionsKilled"`
	VisionScore                 int    `json:"visionScore"`
}

type LeagueRank struct {
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
}

type LeagueProfilePage struct {
	PUUID         string        `json:"PUUID"`
	ProfileIconID int           `json:"profileIconId"`
	GameName      string        `json:"gameName"`
	TagLine       string        `json:"tagLine"`
	Region        string        `json:"region"`
	Level         int           `json:"summonerLevel"`
	Ranks         []LeagueRank  `json:"ranks"`
	MasteryData   MasteryData   `json:"masteryData"`
	MatchData     []LeagueMatch `json:"matchData"`
}

type RiotLeagueMatchResponse struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		EndOfGameResult    string `json:"endOfGameResult"`
		GameCreation       int64  `json:"gameCreation"`
		GameDuration       int    `json:"gameDuration"`
		GameEndTimestamp   int64  `json:"gameEndTimestamp"`
		GameID             int64  `json:"gameId"`
		GameMode           string `json:"gameMode"`
		GameName           string `json:"gameName"`
		GameStartTimestamp int64  `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		GameVersion        string `json:"gameVersion"`
		MapID              int    `json:"mapId"`
		Participants       []struct {
			PlayerScore0  int `json:"PlayerScore0"`
			PlayerScore1  int `json:"PlayerScore1"`
			PlayerScore10 int `json:"PlayerScore10"`
			PlayerScore11 int `json:"PlayerScore11"`
			PlayerScore2  int `json:"PlayerScore2"`
			PlayerScore3  int `json:"PlayerScore3"`
			PlayerScore4  int `json:"PlayerScore4"`
			PlayerScore5  int `json:"PlayerScore5"`
			PlayerScore6  int `json:"PlayerScore6"`
			PlayerScore7  int `json:"PlayerScore7"`
			PlayerScore8  int `json:"PlayerScore8"`
			PlayerScore9  int `json:"PlayerScore9"`
			AllInPings    int `json:"allInPings"`
			AssistMePings int `json:"assistMePings"`
			Assists       int `json:"assists"`
			BaronKills    int `json:"baronKills"`
			BasicPings    int `json:"basicPings"`
			Challenges    struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				HealFromMapSources                       float64 `json:"HealFromMapSources"`
				InfernalScalePickup                      int     `json:"InfernalScalePickup"`
				SWARMDefeatAatrox                        int     `json:"SWARM_DefeatAatrox"`
				SWARMDefeatBriar                         int     `json:"SWARM_DefeatBriar"`
				SWARMDefeatMiniBosses                    int     `json:"SWARM_DefeatMiniBosses"`
				SWARMEvolveWeapon                        int     `json:"SWARM_EvolveWeapon"`
				SWARMHave3Passives                       int     `json:"SWARM_Have3Passives"`
				SWARMKillEnemy                           int     `json:"SWARM_KillEnemy"`
				SWARMPickupGold                          int     `json:"SWARM_PickupGold"`
				SWARMReachLevel50                        int     `json:"SWARM_ReachLevel50"`
				SWARMSurvive15Min                        int     `json:"SWARM_Survive15Min"`
				SWARMWinWith5EvolvedWeapons              int     `json:"SWARM_WinWith5EvolvedWeapons"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                   float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilled                        int     `json:"firstTurretKilled"`
				FirstTurretKilledTime                    float64 `json:"firstTurretKilledTime"`
				FistBumpParticipation                    int     `json:"fistBumpParticipation"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      float64 `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                            int     `json:"killingSprees"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LegendaryItemUsed                        []int   `json:"legendaryItemUsed"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MejaisFullStackInTime                    int     `json:"mejaisFullStackInTime"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				TeleportTakedowns                        int     `json:"teleportTakedowns"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				TwoWardsOneSweeperCount                  int     `json:"twoWardsOneSweeperCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				VoidMonsterKill                          int     `json:"voidMonsterKill"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges"`
			ChampExperience             int    `json:"champExperience"`
			ChampLevel                  int    `json:"champLevel"`
			ChampionID                  int    `json:"championId"`
			ChampionName                string `json:"championName"`
			ChampionTransform           int    `json:"championTransform"`
			CommandPings                int    `json:"commandPings"`
			ConsumablesPurchased        int    `json:"consumablesPurchased"`
			DamageDealtToBuildings      int    `json:"damageDealtToBuildings"`
			DamageDealtToObjectives     int    `json:"damageDealtToObjectives"`
			DamageDealtToTurrets        int    `json:"damageDealtToTurrets"`
			DamageSelfMitigated         int    `json:"damageSelfMitigated"`
			DangerPings                 int    `json:"dangerPings"`
			Deaths                      int    `json:"deaths"`
			DetectorWardsPlaced         int    `json:"detectorWardsPlaced"`
			DoubleKills                 int    `json:"doubleKills"`
			DragonKills                 int    `json:"dragonKills"`
			EligibleForProgression      bool   `json:"eligibleForProgression"`
			EnemyMissingPings           int    `json:"enemyMissingPings"`
			EnemyVisionPings            int    `json:"enemyVisionPings"`
			FirstBloodAssist            bool   `json:"firstBloodAssist"`
			FirstBloodKill              bool   `json:"firstBloodKill"`
			FirstTowerAssist            bool   `json:"firstTowerAssist"`
			FirstTowerKill              bool   `json:"firstTowerKill"`
			GameEndedInEarlySurrender   bool   `json:"gameEndedInEarlySurrender"`
			GameEndedInSurrender        bool   `json:"gameEndedInSurrender"`
			GetBackPings                int    `json:"getBackPings"`
			GoldEarned                  int    `json:"goldEarned"`
			GoldSpent                   int    `json:"goldSpent"`
			HoldPings                   int    `json:"holdPings"`
			IndividualPosition          string `json:"individualPosition"`
			InhibitorKills              int    `json:"inhibitorKills"`
			InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
			InhibitorsLost              int    `json:"inhibitorsLost"`
			Item0                       int    `json:"item0"`
			Item1                       int    `json:"item1"`
			Item2                       int    `json:"item2"`
			Item3                       int    `json:"item3"`
			Item4                       int    `json:"item4"`
			Item5                       int    `json:"item5"`
			Item6                       int    `json:"item6"`
			ItemsPurchased              int    `json:"itemsPurchased"`
			KillingSprees               int    `json:"killingSprees"`
			Kills                       int    `json:"kills"`
			Lane                        string `json:"lane"`
			LargestCriticalStrike       int    `json:"largestCriticalStrike"`
			LargestKillingSpree         int    `json:"largestKillingSpree"`
			LargestMultiKill            int    `json:"largestMultiKill"`
			LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
			MagicDamageDealt            int    `json:"magicDamageDealt"`
			MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
			MagicDamageTaken            int    `json:"magicDamageTaken"`
			Missions                    struct {
				PlayerScore0  int `json:"playerScore0"`
				PlayerScore1  int `json:"playerScore1"`
				PlayerScore2  int `json:"playerScore2"`
				PlayerScore3  int `json:"playerScore3"`
				PlayerScore4  int `json:"playerScore4"`
				PlayerScore5  int `json:"playerScore5"`
				PlayerScore6  int `json:"playerScore6"`
				PlayerScore7  int `json:"playerScore7"`
				PlayerScore8  int `json:"playerScore8"`
				PlayerScore9  int `json:"playerScore9"`
				PlayerScore10 int `json:"playerScore10"`
				PlayerScore11 int `json:"playerScore11"`
			} `json:"missions"`
			NeedVisionPings         int `json:"needVisionPings"`
			NeutralMinionsKilled    int `json:"neutralMinionsKilled"`
			NexusKills              int `json:"nexusKills"`
			NexusLost               int `json:"nexusLost"`
			NexusTakedowns          int `json:"nexusTakedowns"`
			ObjectivesStolen        int `json:"objectivesStolen"`
			ObjectivesStolenAssists int `json:"objectivesStolenAssists"`
			OnMyWayPings            int `json:"onMyWayPings"`
			ParticipantID           int `json:"participantId"`
			PentaKills              int `json:"pentaKills"`
			Perks                   struct {
				StatPerks struct {
					Defense int `json:"defense"`
					Flex    int `json:"flex"`
					Offense int `json:"offense"`
				} `json:"statPerks"`
				Styles []struct {
					Description string `json:"description"`
					Selections  []struct {
						Perk int `json:"perk"`
						Var1 int `json:"var1"`
						Var2 int `json:"var2"`
						Var3 int `json:"var3"`
					} `json:"selections"`
					Style int `json:"style"`
				} `json:"styles"`
			} `json:"perks"`
			PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
			Placement                      int    `json:"placement"`
			PlayerAugment1                 int    `json:"playerAugment1"`
			PlayerAugment2                 int    `json:"playerAugment2"`
			PlayerAugment3                 int    `json:"playerAugment3"`
			PlayerAugment4                 int    `json:"playerAugment4"`
			PlayerAugment5                 int    `json:"playerAugment5"`
			PlayerAugment6                 int    `json:"playerAugment6"`
			PlayerSubteamID                int    `json:"playerSubteamId"`
			ProfileIcon                    int    `json:"profileIcon"`
			PushPings                      int    `json:"pushPings"`
			Puuid                          string `json:"puuid"`
			QuadraKills                    int    `json:"quadraKills"`
			RetreatPings                   int    `json:"retreatPings"`
			RiotIDGameName                 string `json:"riotIdGameName"`
			RiotIDTagline                  string `json:"riotIdTagline"`
			Role                           string `json:"role"`
			SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
			Spell1Casts                    int    `json:"spell1Casts"`
			Spell2Casts                    int    `json:"spell2Casts"`
			Spell3Casts                    int    `json:"spell3Casts"`
			Spell4Casts                    int    `json:"spell4Casts"`
			SubteamPlacement               int    `json:"subteamPlacement"`
			Summoner1Casts                 int    `json:"summoner1Casts"`
			Summoner1ID                    int    `json:"summoner1Id"`
			Summoner2Casts                 int    `json:"summoner2Casts"`
			Summoner2ID                    int    `json:"summoner2Id"`
			SummonerID                     string `json:"summonerId"`
			SummonerLevel                  int    `json:"summonerLevel"`
			SummonerName                   string `json:"summonerName"`
			TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
			TeamID                         int    `json:"teamId"`
			TeamPosition                   string `json:"teamPosition"`
			TimeCCingOthers                int    `json:"timeCCingOthers"`
			TimePlayed                     int    `json:"timePlayed"`
			TotalAllyJungleMinionsKilled   int    `json:"totalAllyJungleMinionsKilled"`
			TotalDamageDealt               int    `json:"totalDamageDealt"`
			TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
			TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
			TotalDamageTaken               int    `json:"totalDamageTaken"`
			TotalEnemyJungleMinionsKilled  int    `json:"totalEnemyJungleMinionsKilled"`
			TotalHeal                      int    `json:"totalHeal"`
			TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
			TotalMinionsKilled             int    `json:"totalMinionsKilled"`
			TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
			TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
			TotalUnitsHealed               int    `json:"totalUnitsHealed"`
			TripleKills                    int    `json:"tripleKills"`
			TrueDamageDealt                int    `json:"trueDamageDealt"`
			TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                int    `json:"trueDamageTaken"`
			TurretKills                    int    `json:"turretKills"`
			TurretTakedowns                int    `json:"turretTakedowns"`
			TurretsLost                    int    `json:"turretsLost"`
			UnrealKills                    int    `json:"unrealKills"`
			VisionClearedPings             int    `json:"visionClearedPings"`
			VisionScore                    int    `json:"visionScore"`
			VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
			WardsKilled                    int    `json:"wardsKilled"`
			WardsPlaced                    int    `json:"wardsPlaced"`
			Win                            bool   `json:"win"`
		} `json:"participants"`
		PlatformID string `json:"platformId"`
		QueueID    int    `json:"queueId"`
		Teams      []struct {
			Bans []struct {
				ChampionID int `json:"championId"`
				PickTurn   int `json:"pickTurn"`
			} `json:"bans"`
			Feats struct {
				EPICMONSTERKILL struct {
					FeatState int `json:"featState"`
				} `json:"EPIC_MONSTER_KILL"`
				FIRSTBLOOD struct {
					FeatState int `json:"featState"`
				} `json:"FIRST_BLOOD"`
				FIRSTTURRET struct {
					FeatState int `json:"featState"`
				} `json:"FIRST_TURRET"`
			} `json:"feats"`
			Objectives struct {
				Atakhan struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"atakhan"`
				Baron struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"baron"`
				Champion struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"champion"`
				Dragon struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"dragon"`
				Horde struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"horde"`
				Inhibitor struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"inhibitor"`
				RiftHerald struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"riftHerald"`
				Tower struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"tower"`
			} `json:"objectives"`
			TeamID int  `json:"teamId"`
			Win    bool `json:"win"`
		} `json:"teams"`
		TournamentCode string `json:"tournamentCode"`
	} `json:"info"`
}
