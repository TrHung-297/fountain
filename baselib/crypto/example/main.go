package main

import (
	"encoding/json"
	"log"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/crypto"
)

type GameID string
type GameState string
type StatisticsReq struct {
	GameID GameID `json:"game_id,omitempty"`

	// Playing
	GoldCollected    int32 `json:"gold_collected,omitempty"`
	VillagerHigh     int32 `json:"villager_high,omitempty"`
	TotalPopulation  int32 `json:"total_population,omitempty"`
	LivingPopulation int32 `json:"living_population,omitempty"`
	Exploration      int32 `json:"exploration,omitempty"`
	TributeGiven     int32 `json:"tribute_given,omitempty"` // Chỉ số bơm đồ
	Kills            int32 `json:"kills,omitempty"`
	Razings          int32 `json:"razings,omitempty"`
	Losses           int32 `json:"losses,omitempty"`
	Technologies     int32 `json:"technologies,omitempty"`
	Age              int8  `json:"age,omitempty"`
	AdvancedTime     int32 `json:"advanced_time,omitempty"` // In seconds

	// Ending
	Result                int8  `json:"result,omitempty"`
	StoneAgeUpgradedTime  int32 `json:"stone_age_upgraded_time,omitempty"`  // In seconds
	BronzeAgeUpgradedTime int32 `json:"bronze_age_upgraded_time,omitempty"` // In seconds
	SteelAgeUpgradedTime  int32 `json:"steel_age_upgraded_time,omitempty"`

	Status GameState `json:"status,omitempty"`

	CreatedTime int32 `json:"created_time,omitempty"`
}

func main() {
	cryptor := crypto.NewAES256IGECryptor([]byte("scPQ4l2lHNZpWQwn8vgH0QUr4iQw9mHr"), []byte("jKPOOoiAW2MfNxRN"))
	msg := "haha hihi tao ne"
	dataEn, err := cryptor.Encrypt([]byte(msg))
	if err != nil {
		panic(err)
	}
	log.Printf("Data En: %+v", dataEn)

	dataDe, err := cryptor.Decrypt(dataEn)
	if err != nil {
		panic(err)
	}
	log.Printf("Data De: %s", dataDe)

	base64Encrypted := `6KInEH0GwemShWg3syA4UigM6a8PTJRxQJuFdC1OAh+sYlx8HAZARTqGYhRVHzEwkF9y3+iPP1E0IUfpV33402r+9ooEGCGKEXCj9puT9FN+ZXfAsmIsZOLpMyg7VEuc`
	data, err := cryptor.DecryptSimpleWithBase64(base64Encrypted)
	if err != nil {
		panic(err)
	}
	log.Printf("DecryptSimpleWithBase64 Data De: %s", data)

	req := &StatisticsReq{}
	if err := json.Unmarshal(data, req); err != nil {
		log.Printf("RankingStatisticsAPI::Save - BodyParser Error: %+v", err)

		panic(err)
	}
	log.Printf("RankingStatisticsAPI::Save - Request: %s", base.JSONDebugDataString(req))
}
