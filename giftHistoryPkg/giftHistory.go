package giftHistoryPkg

import (
	"encoding/json"
	"errors"
	"github.com/eatobin/redpoint-go/giftPairPkg"
)

type History []giftPairPkg.StructGiftPair

// JsonStringToGiftHistory turns a JSON string into a GiftHistory
func JsonStringToGiftHistory(jsonString string) (History, error) {
	var giftHistory History
	err := json.Unmarshal([]byte(jsonString), &giftHistory)
	if err != nil {
		return History{}, err
	}
	for _, v := range giftHistory {
		if v.Givee == "" || v.Giver == "" {
			err = errors.New("missing one or both field values somewhere")
			return History{}, err
		}
	}
	return giftHistory, nil
}

// AddYear adds a playerKey to each giftHistory
func (giftHistory History) AddYear(playerKey string) History {
	return append(giftHistory, giftPairPkg.StructGiftPair{Givee: playerKey, Giver: playerKey})
}

// UpdateGiftHistory replaces a giftPair at a given giftYear
func (giftHistory History) UpdateGiftHistory(giftYear int, giftPair giftPairPkg.StructGiftPair) History {
	newGiftHistory := make(History, len(giftHistory))
	copy(newGiftHistory, giftHistory)
	newGiftHistory[giftYear] = giftPair
	return newGiftHistory
}
