package giftHistory

import (
	"encoding/json"
	"errors"
	"github.com/eatobin/redpoint-go/giftPair"
)

type HistoryTA = []giftPair.Struct
type GiftYearTA = int

// HistoryAssertEqual compares two GiftHistories
func HistoryAssertEqual(a, b HistoryTA) bool {
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

// JsonStringToGiftHistory turns a JSON string into a GiftHistory
func JsonStringToGiftHistory(jsonString giftPair.JsonStringTA) (HistoryTA, error) {
	var giftHistory HistoryTA
	err := json.Unmarshal([]byte(jsonString), &giftHistory)
	if err != nil {
		return HistoryTA{}, err
	}
	for _, v := range giftHistory {
		if v.Givee == "" || v.Giver == "" {
			err = errors.New("missing one or both field values somewhere")
			return HistoryTA{}, err
		}
	}
	return giftHistory, nil
}

// AddYear adds a playerKey to each giftHistory
func AddYear(playerKey giftPair.PlayerKeyTA, giftHistory HistoryTA) HistoryTA {
	return append(giftHistory, giftPair.Struct{Givee: playerKey, Giver: playerKey})
}

// UpdateGiftHistory replaces a giftPair at a given giftYear
func UpdateGiftHistory(giftYear GiftYearTA, giftPair giftPair.Struct, giftHistory HistoryTA) HistoryTA {
	newGiftHistory := make(HistoryTA, len(giftHistory))
	copy(newGiftHistory, giftHistory)
	newGiftHistory[giftYear] = giftPair
	return newGiftHistory
}
