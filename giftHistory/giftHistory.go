package giftHistory

import (
	"encoding/json"
	"github.com/eatobin/redpointGo/giftPair"
)

type GiftPair = giftPair.GiftPair
type GiftHistory = []GiftPair

// GHcompare tells whether GHa and GHb contain the same elements.
// A nil argument is equivalent to an empty slice.
func GHcompare(a, b []GiftPair) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// AddYear adds a playerKey to each giftHistory
func AddYear(playerKey string, giftHistory GiftHistory) GiftHistory {
	return append(giftHistory, GiftPair{Givee: playerKey, Giver: playerKey})
}

// UpdateGiftHistory replaces a giftPair at a given giftYear
func UpdateGiftHistory(giftYear int, giftPair GiftPair, giftHistory GiftHistory) GiftHistory {
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
