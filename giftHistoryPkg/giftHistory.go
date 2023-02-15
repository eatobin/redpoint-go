package giftHistoryPkg

import (
	"encoding/json"
	"github.com/eatobin/redpoint-go/giftPairPkg"
)

type GiftHistory = []giftPairPkg.GiftPair

// GiftHistoryAssertEqual compares two GiftHistories
func GiftHistoryAssertEqual(a, b GiftHistory) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !giftPairPkg.GiftPairAssertEqual(v, b[i]) {
			return false
		}
	}
	return true
}

// AddYear adds a playerKey to each giftHistory
func AddYear(playerKey string, giftHistory GiftHistory) GiftHistory {
	return append(giftHistory, giftPairPkg.GiftPair{Givee: playerKey, Giver: playerKey})
}

// UpdateGiftHistory replaces a giftPair at a given giftYear
func UpdateGiftHistory(giftYear int, giftPair giftPairPkg.GiftPair, giftHistory GiftHistory) GiftHistory {
	giftHistory[giftYear] = giftPair
	return giftHistory
}

// JsonStringToGiftHistory turns a GiftHistory JSON string into a GiftHistory
func JsonStringToGiftHistory(ghString string) (GiftHistory, error) {
	var giftHistory GiftHistory
	err := json.Unmarshal([]byte(ghString), &giftHistory)
	return giftHistory, err
}

// GHToJsonString turns a GiftHistory into a GiftHistory JSON string
func GHToJsonString(giftHistory GiftHistory) (string, error) {
	ghByte, err := json.Marshal(giftHistory)
	return string(ghByte), err
}
