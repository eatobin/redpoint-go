package players

import (
	"encoding/json"
	"github.com/eatobin/redpoint-go/giftHistoryPkg"
	"github.com/eatobin/redpoint-go/giftPairPkg"
	"github.com/eatobin/redpoint-go/player"
)

type Givee = giftPairPkg.GiveeTA
type Giver = giftPairPkg.GiverTA
type GiftPair = giftPairPkg.GiftPair
type GiftHistory = giftHistoryPkg.GiftHistoryTA
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
		ngh := giftHistoryPkg.AddYear(playerKey, gh)
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
	ngh := giftHistoryPkg.UpdateGiftHistory(giftYear, giftPair, players[playerKey].GiftHistory)
	nplr := players[playerKey].UpdateGiftHistory(ngh)
	return UpdatePlayer(playerKey, nplr, players)
}

func UpdateGivee(selfKey string, giftYear int, givee Givee, players Players) Players {
	ngp := players[selfKey].GiftHistory[giftYear].GiftPairUpdateGivee(givee)
	return setGiftPair(selfKey, giftYear, ngp, players)
}

func UpdateGiver(selfKey string, giftYear int, giver Giver, players Players) Players {
	ngp := players[selfKey].GiftHistory[giftYear].GiftPairUpdateGiver(giver)
	return setGiftPair(selfKey, giftYear, ngp, players)
}

// JsonStringToPlayers turns a Players JSON string into a Players
func JsonStringToPlayers(ghString string) (Players, error) {
	var players Players
	err := json.Unmarshal([]byte(ghString), &players)
	return players, err
}
