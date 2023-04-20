package main

import (
	"encoding/json"
	"errors"
)

type History []GiftPairStruct

// GiftHistoryJsonStringToGiftHistory turns a JSON string into a GiftHistory
func GiftHistoryJsonStringToGiftHistory(jsonString string) (History, error) {
	var giftHistory History
	err := json.Unmarshal([]byte(jsonString), &giftHistory)
	if err != nil {
		return History{}, err
	}
	for _, v := range giftHistory {
		if v.Givee == "" || v.Giver == "" {
			err = errors.New("missing one or both GiftPair field values somewhere")
			return History{}, err
		}
	}
	return giftHistory, nil
}

// GiftHistoryAddYear adds a playerKey to each giftHistory
func (giftHistory History) GiftHistoryAddYear(playerKey string) History {
	return append(giftHistory, GiftPairStruct{Givee: playerKey, Giver: playerKey})
}

// GiftHistoryUpdateGiftHistory replaces a giftPair at a given giftYear
func (giftHistory History) GiftHistoryUpdateGiftHistory(giftYear int, giftPair GiftPairStruct) History {
	newGiftHistory := make(History, len(giftHistory))
	copy(newGiftHistory, giftHistory)
	newGiftHistory[giftYear] = giftPair
	return newGiftHistory
}
