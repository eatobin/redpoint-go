package players

import (
	"encoding/json"
	"github.com/eatobin/redpoint-go/giftHistory"
	"github.com/eatobin/redpoint-go/giftPair"
	"github.com/eatobin/redpoint-go/player"
)

type Givee = giftPair.Givee
type Giver = giftPair.Giver
type GiftPair = giftPair.GiftPair
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

func GetGivee(selfKey string, giftYear int, players Players) Givee {
	return players[selfKey].GiftHistory[giftYear].Givee
}

func GetGiver(selfKey string, giftYear int, players Players) Giver {
	return players[selfKey].GiftHistory[giftYear].Giver
}

func setGiftPair(playerKey string, giftYear int, giftPair GiftPair, players Players) Players {
	ngh := giftHistory.UpdateGiftHistory(giftYear, giftPair, players[playerKey].GiftHistory)
	nplr := players[playerKey].UpdateGiftHistory(ngh)
	return UpdatePlayer(playerKey, nplr, players)
}

func UpdateGivee(selfKey string, giftYear int, givee Givee, players Players) Players {
	ngp := players[selfKey].GiftHistory[giftYear].UpdateGivee(givee)
	return setGiftPair(selfKey, giftYear, ngp, players)
}

func UpdateGiver(selfKey string, giftYear int, giver Giver, players Players) Players {
	ngp := players[selfKey].GiftHistory[giftYear].UpdateGiver(giver)
	return setGiftPair(selfKey, giftYear, ngp, players)
}

// JsonStringToPlayers turns a Players JSON string into a Players
func JsonStringToPlayers(ghString string) (Players, error) {
	var players Players
	err := json.Unmarshal([]byte(ghString), &players)
	return players, err
}
