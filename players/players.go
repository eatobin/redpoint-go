package players

import (
	"github.com/eatobin/redpointGo/giftHistory"
	"github.com/eatobin/redpointGo/player"
)

//type GiftPair = giftPair.GiftPair
type GiftHistory = giftHistory.GiftHistory
type Player = player.Player
type Players = map[string]Player

// ComparePlayers tells whether PlayersA and PlayersB contain the same elements.
// A nil argument is equivalent to an empty slice.
func ComparePlayers(a, b Players) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if player.ComparePlayer(v, b[k]) {
			return false
		}
	}
	return true
}

func UpdatePlayer(playerKey string, player Player, players Players) Players {
	players[playerKey] = player
	return players
}

func GetPlayerName(playerKey string, players Players) string {
	return players[playerKey].PlayerName
}
