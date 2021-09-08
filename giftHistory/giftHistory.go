package giftHistory

import "eatobin.com/redpointGo/giftPair"

type GiftPair = giftPair.GiftPair
type GiftHistory = []GiftPair

func addYear(playerKey string, giftHistory GiftHistory) GiftHistory {
	return append(giftHistory, GiftPair{Givee: playerKey, Giver: playerKey})
}
