package giftHistoryPkg

import (
	"encoding/json"
	"github.com/eatobin/redpoint-go/giftPairPkg"
)

type GiftHistoryTA = []giftPairPkg.GiftPair
type GiftYearTA = int

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

// GiftHistoryJsonStringToGiftHistory turns a JSON string into a GiftHistory
func GiftHistoryJsonStringToGiftHistory(jsonString giftPairPkg.JsonStringTA) (GiftHistoryTA, error) {
	var giftHistory GiftHistoryTA
	err := json.Unmarshal([]byte(jsonString), &giftHistory)
	return giftHistory, err
}

// GiftHistoryAddYear adds a playerKey to each giftHistory
func GiftHistoryAddYear(playerKey giftPairPkg.PlayerKeyTA, giftHistory GiftHistoryTA) GiftHistoryTA {
	return append(giftHistory, giftPairPkg.GiftPair{Givee: playerKey, Giver: playerKey})
}

// GiftHistoryUpdateGiftHistory replaces a giftPair at a given giftYear
func GiftHistoryUpdateGiftHistory(giftYear GiftYearTA, giftPair giftPairPkg.GiftPair, giftHistory GiftHistoryTA) GiftHistoryTA {
	giftHistory[giftYear] = giftPair
	return giftHistory
}
