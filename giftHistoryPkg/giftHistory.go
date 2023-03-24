package giftHistoryPkg

import (
	"encoding/json"
	"github.com/eatobin/redpoint-go/giftPairPkg"
)

type GiftHistoryTA = []giftPairPkg.GiftPair

// GiftHistoryAssertEqual compares two GiftHistories
func GiftHistoryAssertEqual(a, b GiftHistoryTA) bool {
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
func AddYear(playerKey string, giftHistory GiftHistoryTA) GiftHistoryTA {
	return append(giftHistory, giftPairPkg.GiftPair{Givee: playerKey, Giver: playerKey})
}

// UpdateGiftHistory replaces a giftPair at a given giftYear
func UpdateGiftHistory(giftYear int, giftPair giftPairPkg.GiftPair, giftHistory GiftHistoryTA) GiftHistoryTA {
	giftHistory[giftYear] = giftPair
	return giftHistory
}

// JsonStringToGiftHistory turns a GiftHistory JSON string into a GiftHistory
func JsonStringToGiftHistory(ghString string) (GiftHistoryTA, error) {
	var giftHistory GiftHistoryTA
	err := json.Unmarshal([]byte(ghString), &giftHistory)
	return giftHistory, err
}

// GHToJsonString turns a GiftHistory into a GiftHistory JSON string
func GHToJsonString(giftHistory GiftHistoryTA) (string, error) {
	ghByte, err := json.Marshal(giftHistory)
	return string(ghByte), err
}
