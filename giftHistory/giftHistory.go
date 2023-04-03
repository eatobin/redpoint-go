package giftHistory

import (
	"encoding/json"
	"errors"
	"github.com/eatobin/redpoint-go/giftPair"
)

type History = []giftPair.Struct
type GiftYear = int

// JsonStringToGiftHistory turns a JSON string into a GiftHistory
func JsonStringToGiftHistory(jsonString giftPair.JsonString) (History, error) {
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
func AddYear(playerKey giftPair.PlayerKey, giftHistory History) History {
	return append(giftHistory, giftPair.Struct{Givee: playerKey, Giver: playerKey})
}

// UpdateGiftHistory replaces a giftPair at a given giftYear
func UpdateGiftHistory(giftYear GiftYear, giftPair giftPair.Struct, giftHistory History) History {
	newGiftHistory := make(History, len(giftHistory))
	copy(newGiftHistory, giftHistory)
	newGiftHistory[giftYear] = giftPair
	return newGiftHistory
}
