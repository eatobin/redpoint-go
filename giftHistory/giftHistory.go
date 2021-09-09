package giftHistory

import "github.com/eatobin/redpointGo/giftPair"

type GiftPair = giftPair.GiftPair
type GiftHistory = []GiftPair

func AddYear(playerKey string, giftHistory GiftHistory) GiftHistory {
	return append(giftHistory, GiftPair{Givee: playerKey, Giver: playerKey})
}

func UpdateGiftHistory(giftYear int, giftPair GiftPair, giftHistory GiftHistory) GiftHistory {
	giftHistory[giftYear] = giftPair
	return giftHistory
}
