package player

import "testing"

//"A Player" should "return an updated giftHistory" in {
//assert(Player.updateGiftHistory(Vector(GiftPair("nope", "yup")))(player) ==
//Player("Paul McCartney", Vector(GiftPair("nope", "yup"))))
//}
//
//it should "convert from JSON" in {
//val plrJson: Either[Error, Player] = Player.jsonStringToPlayer(jsonStringPlr)
//assert(plrJson == Right(player))
//}
//
//it should "convert to JSON" in {
//val plrJson: JsonString = Player.playerToJsonString(player)
//assert(plrJson == jsonStringPlr)
//}
//}

var jsonStringPlr = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var player = Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}}

func TestUpdateGiftHistory(t *testing.T) {
	newGH := GiftHistory{{Givee: "nope", Giver: "yup"}}
	newPlayer := Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "nope", Giver: "yup"}}}.String()
	gotPlayer := player.UpdateGiftHistory(newGH).String()
	if gotPlayer != newPlayer {
		t.Fatalf("(%v) UpdateGiftHistory(%v) ==\n%v,\nwant %v", player, newGH, gotPlayer, newPlayer)
	}
}

//func TestJsonStringToGiftPair(t *testing.T) {
//	gotGP, _ := JsonStringToGiftPair(jsonStringGP)
//	if gotGP != gp {
//		t.Fatalf("JsonStringToGiftPair(%v) == %v, want %v", jsonStringGP, gotGP, gp)
//	}
//}
//
//func TestGiftPairToJsonString(t *testing.T) {
//	gotString, _ := gp.GiftPairToJsonString()
//	if gotString != jsonStringGP {
//		t.Fatalf("(%v) JsonStringToGiftPair() == %v, want %v", gp, gotString, jsonStringGP)
//	}
//}
