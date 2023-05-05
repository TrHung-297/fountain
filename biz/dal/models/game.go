

package models

type GameState string
type GameID string // Google UUID
type GameName string
type GameMode int
type GameType int
type GameKind int
type GameRoundType int
type RoomStatus int

const (
	KGameStateOpening       GameState = "opening"    // Only for party
	KGameStateWaiting       GameState = "waiting"    // Both for game and party
	KGameStateFoundGame     GameState = "found"      // Only for party
	KGameStateReady         GameState = "ready"      // Both for game and party
	KGameStateGaming        GameState = "gaming"     // Both for game and party
	KGameStatePrepareFinish GameState = "pre_finish" // Both for game and party
	KGameStateFinish        GameState = "finish"     // Extend for game, party also can use too
	KGameStateReject        GameState = "reject"     // Extend for game, party also can use too
	KGameStateTimeout       GameState = "timeout"    // Extend for game, party also can use too

	KGameStateName = "game_state"
	KGameIDName    = "game_id"
	KGameNameName  = "game_name"

	KGameRoundTypeName                                     = "game_event_type"
	KGameEventRoundEliminationTournamentType GameRoundType = 0
	KGameEventRoundRobinTournamentType       GameRoundType = 1

	KGameModeName          = "game_mode"
	KGameMode1vs1 GameMode = 0
	KGameMode2vs2 GameMode = 1
	KGameMode3vs3 GameMode = 2
	KGameMode4vs4 GameMode = 3

	KGameTypeName                  = "game_type"
	KGameTypeRandom       GameType = 0
	KGameTypeRBowShang    GameType = 1
	KGameTypeRBowAssyrian GameType = 2
	KGameDeathMatch       GameType = 3
	KGameTypeAllShang     GameType = 4
	KGameTypeAllAssyrian  GameType = 5

	KGameOfflineRoomStatusName        = "game_room_status"
	KGameOfflineRoomStatusInit    int = 0
	KGameOfflineRoomStatusFull    int = 1
	KGameOfflineRoomStatusRunning int = 2
	KGameOfflineRoomStatusSuccess int = 3
	KGameOfflineRoomStatusFailled int = -1

	KGameKindName              = "game_kind"
	KGameKindCustom   GameKind = 0
	KGameAutoMatching GameKind = 1
	KGameTournament   GameKind = 6
	KGameRanking      GameKind = 7
	KGameArena        GameKind = 8
)
