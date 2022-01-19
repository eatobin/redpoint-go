package players

import (
	"github.com/eatobin/redpoint-go/giftHistory"
	"github.com/eatobin/redpoint-go/player"
)

//type GiftPair = giftPair.GiftPair
type GiftHistory = giftHistory.GiftHistory
type Player = player.Player
type Players = map[string]Player

func ComparePlayers(a, b Players) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if !player.ComparePlayer(v, b[k]) {
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

func AddYear(players Players) Players {
	for playerKey, thisPlayer := range players {
		gh := thisPlayer.GiftHistory
		ngh := giftHistory.AddYear(playerKey, gh)
		nplr := thisPlayer.UpdateGiftHistory(ngh)
		players[playerKey] = nplr
	}
	return players
}
