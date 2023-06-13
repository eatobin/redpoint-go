package main

import (
	"encoding/json"
	"errors"
)

type Players map[string]PlayerStruct

// PlayersJsonStringToPlayers turns a JSON string into a Players
func PlayersJsonStringToPlayers(jsonString string) (Players, error) {
	var players Players
	err := json.Unmarshal([]byte(jsonString), &players)
	if err != nil {
		return Players{}, err
	}
	for k := range players {
		if k == "" {
			err = errors.New("missing PlayerKey")
			return Players{}, err
		}
	}
	for _, v := range players {
		if v.PlayerName == "" {
			err = errors.New("missing PlayerName field value")
			return Players{}, err
		}
		if len(v.GiftHistory) == 0 {
			err = errors.New("missing GiftHistory field value")
			return Players{}, err
		}
	}
	for _, v := range players {
		for _, gp := range v.GiftHistory {
			if gp.Givee == "" || gp.Giver == "" {
				err = errors.New("missing one or both GiftPair field values somewhere")
				return Players{}, err
			}
		}
	}
	return players, nil
}

func (players Players) PlayersUpdatePlayer(playerKey string, player PlayerStruct) Players {
	playersCopy := make(Players)
	for k, v := range players {
		playersCopy[k] = v
	}
	playersCopy[playerKey] = player
	return playersCopy
}

func (players Players) PlayersGetPlayerName(playerKey string) string {
	return players[playerKey].PlayerName
}

func (players Players) PlayersAddYear() Players {
	playersCopy := make(Players)
	for k, v := range players {
		playersCopy[k] = v
	}
	for playerKey, player := range playersCopy {
		gh := player.GiftHistory
		ngh := gh.GiftHistoryAddYear(playerKey)
		nplr := player.PlayerUpdateGiftHistory(ngh)
		playersCopy[playerKey] = nplr
	}
	return playersCopy
}

// func GetGivee(selfKey string, giftYear int, players Players) Givee {
// 	return players[selfKey].GiftHistory[giftYear].Givee
// }

// func GetGiver(selfKey string, giftYear int, players Players) Giver {
// 	return players[selfKey].GiftHistory[giftYear].Giver
// }

// func setGiftPair(playerKey string, giftYear int, giftPair GiftPair, players Players) Players {
// 	ngh := giftHistory.UpdateGiftHistory(giftYear, giftPair, players[playerKey].GiftHistory)
// 	nplr := players[playerKey].UpdateGiftHistory(ngh)
// 	return UpdatePlayer(playerKey, nplr, players)
// }

// func UpdateGivee(selfKey string, giftYear int, givee Givee, players Players) Players {
// 	ngp := players[selfKey].GiftHistory[giftYear].GiftPairUpdateGivee(givee)
// 	return setGiftPair(selfKey, giftYear, ngp, players)
// }

// func UpdateGiver(selfKey string, giftYear int, giver Giver, players Players) Players {
// 	ngp := players[selfKey].GiftHistory[giftYear].GiftPairUpdateGiver(giver)
// 	return setGiftPair(selfKey, giftYear, ngp, players)
// }
