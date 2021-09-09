package giftHistory

import (
	"encoding/json"
	"github.com/eatobin/redpointGo/giftPair"
)

type GiftPair = giftPair.GiftPair
type GiftPairPtr = *GiftPair
type GiftHistory = *[]GiftPairPtr

func AddYear(playerKey string, giftHistory GiftHistory) {
	*giftHistory = append(*giftHistory, &GiftPair{Givee: playerKey, Giver: playerKey})
}

func UpdateGiftHistory(giftYear int, giftPairPtr GiftPairPtr, giftHistory GiftHistory) {
	derefGH := *giftHistory
	derefGH[giftYear] = giftPairPtr
}

// JsonStringToGiftHistory turns a GiftHistory JSON string into a GiftHistory
func JsonStringToGiftHistory(ghString string) (GiftHistory, error) {
	var giftHistory GiftHistory
	err := json.Unmarshal([]byte(ghString), &giftHistory)
	return giftHistory, err
}

// GiftHistoryToJsonString turns a GiftHistory into a GiftHistory JSON string
func GiftHistoryToJsonString(giftHistory GiftHistory) (string, error) {
	ghByte, err := json.Marshal(*giftHistory)
	return string(ghByte), err
}
