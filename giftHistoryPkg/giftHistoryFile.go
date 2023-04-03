package giftHistoryPkg

import (
	"encoding/json"
	"errors"
)

type GiftHistoryTA = []giftPair.GiftPairStruct
type GiftYearTA = int

// GiftHistoryAssertEqual compares two GiftHistories
func GiftHistoryAssertEqual(a, b GiftHistoryTA) bool {
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

// GiftHistoryJsonStringToGiftHistory turns a JSON string into a GiftHistory
func GiftHistoryJsonStringToGiftHistory(jsonString giftPair.JsonStringTA) (GiftHistoryTA, error) {
	var giftHistory GiftHistoryTA
	err := json.Unmarshal([]byte(jsonString), &giftHistory)
	if err != nil {
		return GiftHistoryTA{}, err
	}
	for _, v := range giftHistory {
		if v.Givee == "" || v.Giver == "" {
			err = errors.New("missing one or both field values somewhere")
			return GiftHistoryTA{}, err
		}
	}
	return giftHistory, nil
}

// GiftHistoryAddYear adds a playerKey to each giftHistory
func GiftHistoryAddYear(playerKey giftPair.PlayerKeyTA, giftHistory GiftHistoryTA) GiftHistoryTA {
	return append(giftHistory, giftPair.GiftPairStruct{Givee: playerKey, Giver: playerKey})
}

// GiftHistoryUpdateGiftHistory replaces a giftPair at a given giftYear
func GiftHistoryUpdateGiftHistory(giftYear GiftYearTA, giftPair giftPair.GiftPairStruct, giftHistory GiftHistoryTA) GiftHistoryTA {
	newGiftHistory := make(GiftHistoryTA, len(giftHistory))
	copy(newGiftHistory, giftHistory)
	newGiftHistory[giftYear] = giftPair
	return newGiftHistory
}
