package player

import (
	"github.com/eatobin/redpointGo/giftHistory"
	"github.com/eatobin/redpointGo/giftPair"
)

type GiftPair = giftPair.GiftPair
type GiftHistory = giftHistory.GiftHistory

//case class Player(playerName: String, giftHistory: Vector[GiftPair])
//
//object Player {
//def updateGiftHistory(giftHistory: Vector[GiftPair])(player: Player): Player =
//player.copy(giftHistory = giftHistory)
//
//def jsonStringToPlayer(plrString: String): Either[Error, Player] =
//decode[Player](plrString)
//
//def playerToJsonString(player: Player): JsonString =
//player.asJson.noSpaces
//}

// A Player has a PlayerName and a GiftHistory
type Player struct {
	PlayerName  string `json:"playerName"`
	GiftHistory string `json:"giftHistory"`
}

//// UpdateGivee updates a Givee in a GiftPair
//func (gp GiftPair) UpdateGivee(givee string) GiftPair {
//	gp.Givee = givee
//	return gp
//}
//
//// UpdateGiver updates a Giver in a GiftPair
//func (gp GiftPair) UpdateGiver(giver string) GiftPair {
//	gp.Giver = giver
//	return gp
//}
//
//// String makes a GiftPair into a string
//func (gp GiftPair) String() string {
//	return fmt.Sprintf("{Givee: %s, Giver: %s}", gp.Givee, gp.Giver)
//}
//
//// JsonStringToGiftPair turns a GiftPair JSON string into a GiftPair
//func JsonStringToGiftPair(gpString string) (GiftPair, error) {
//	var giftPair GiftPair
//	err := json.Unmarshal([]byte(gpString), &giftPair)
//	return giftPair, err
//}
//
//// GiftPairToJsonString turns a GiftPair into a GiftPair JSON string
//func (gp GiftPair) GiftPairToJsonString() (string, error) {
//	gpByte, err := json.Marshal(gp)
//	return string(gpByte), err
//}
