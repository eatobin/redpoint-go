package players

import (
	"github.com/eatobin/redpointGo/giftHistory"
	"github.com/eatobin/redpointGo/giftPair"
	"github.com/eatobin/redpointGo/player"
)

type GiftPair = giftPair.GiftPair
type GiftHistory = giftHistory.GiftHistory
type Player = player.Player
type Players = map[string]Player

func UpdatePlayer(playerKey string, player Player, players Players) Players {
	players[playerKey] = player
	return players
}

func GetPlayerName(playerKey string, players Players) string {
	return players[playerKey].PlayerName
}
