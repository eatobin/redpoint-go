package player

import "testing"

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

func TestJsonStringToPlayer(t *testing.T) {
	gotPlayer, _ := JsonStringToPlayer(jsonStringPlr)
	if gotPlayer.String() != player.String() {
		t.Fatalf("JsonStringToGiftPlayer(%v) ==\n%v,\nwant %v", jsonStringPlr, gotPlayer.String(), player.String())
	}
}

func TestPlayerToJsonString(t *testing.T) {
	gotString, _ := player.PlayerToJsonString()
	if gotString != jsonStringPlr {
		t.Fatalf("(%v) JsonStringToPlayer() ==\n%v,\nwant %v", player, gotString, jsonStringPlr)
	}
}
