package giftHistory

import "github.com/eatobin/redpointGo/giftPair"

type GiftPair = giftPair.GiftPair
type GiftPairPtr = *GiftPair
type GiftHistory = []GiftPairPtr

func AddYear(playerKey string, giftHistory GiftHistory) GiftHistory {
	return append(giftHistory, &GiftPair{Givee: playerKey, Giver: playerKey})
}

func UpdateGiftHistory(giftYear int, giftPairPtr GiftPairPtr, giftHistory GiftHistory) {
	giftHistory[giftYear] = giftPairPtr
}
