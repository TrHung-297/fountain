
package models

type SessionID string // Google UUID
type EventID string   // Google UUID
type BetID string     // Google UUID
type SponsorID string // Google UUID

const (
	KSessionIDName = "session_id"
	KEventIDName   = "event_id"
	KBetIDName     = "bet_id"
	KSponsorIDName = "sponsor_id"
)
