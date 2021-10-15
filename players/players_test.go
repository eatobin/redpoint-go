package players

import (
	"testing"
)

//import "testing"

var jsonStringPlrs = "{\"PauMcc\":{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
var rinSta = Player{PlayerName: "Ringo Starr", GiftHistory: GiftHistory{{Givee: "JohLen", Giver: "GeoHar"}}}
var johLen = Player{PlayerName: "John Lennon", GiftHistory: GiftHistory{{Givee: "PauMcc", Giver: "RinSta"}}}
var geoHar = Player{PlayerName: "George Harrison", GiftHistory: GiftHistory{{Givee: "RinSta", Giver: "PauMcc"}}}
var pauMcc = Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}}
var players = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}

var newBee = Player{PlayerName: "New Bee", GiftHistory: GiftHistory{{Givee: "NewBee", Giver: "NewBee"}}}
var newBeePlayers = Players{"RinSta": newBee, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}

//private val rinStaExt: Player = Player("Ringo Starr", Vector(GiftPair("JohLen", "GeoHar"), GiftPair("RinSta", "RinSta")))
//private val johLenExt: Player = Player("John Lennon", Vector(GiftPair("PauMcc", "RinSta"), GiftPair("JohLen", "JohLen")))
//private val geoHarExt: Player = Player("George Harrison", Vector(GiftPair("RinSta", "PauMcc"), GiftPair("GeoHar", "GeoHar")))
//private val pauMccExt: Player = Player("Paul McCartney", Vector(GiftPair("GeoHar", "JohLen"), GiftPair("PauMcc", "PauMcc")))
//private val playersExt: Map[String, Player] =
//Map("RinSta" -> rinStaExt, "JohLen" -> johLenExt, "GeoHar" -> geoHarExt, "PauMcc" -> pauMccExt)
//
//private val geoHarGivee: Player = Player("George Harrison", Vector(GiftPair("you", "PauMcc")))
//private val geoHarGiver: Player = Player("George Harrison", Vector(GiftPair("RinSta", "you")))
//private val playersGivee: Map[String, Player] =
//Map("RinSta" -> rinSta, "JohLen" -> johLen, "GeoHar" -> geoHarGivee, "PauMcc" -> pauMcc)
//private val playersGiver: Map[String, Player] =
//Map("RinSta" -> rinSta, "JohLen" -> johLen, "GeoHar" -> geoHarGiver, "PauMcc" -> pauMcc)

//var jsonStringPlr = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
//var player = Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}}

func TestUpdatePlayer(t *testing.T) {
	goodBee := Player{PlayerName: "New Bee", GiftHistory: GiftHistory{{Givee: "NewBee", Giver: "NewBeeX"}}}
	gotPlayers := UpdatePlayer("RinSta", goodBee, players)
	if ComparePlayers(gotPlayers, newBeePlayers) {
		t.Fatalf("UpdatePlayer(%s, %v, %v) ==\n%v,\nwant %v", "RinSta", newBee, players, gotPlayers, goodBee)
	}
}

//
//func TestJsonStringToPlayer(t *testing.T) {
//	gotPlayer, _ := JsonStringToPlayer(jsonStringPlr)
//	if gotPlayer.String() != player.String() {
//		t.Fatalf("JsonStringToGiftPlayer(%v) ==\n%v,\nwant %v", jsonStringPlr, gotPlayer.String(), player.String())
//	}
//}
//
//func TestPlayerToJsonString(t *testing.T) {
//	gotString, _ := player.PlayerToJsonString()
//	if gotString != jsonStringPlr {
//		t.Fatalf("(%v) JsonStringToPlayer() ==\n%v,\nwant %v", player, gotString, jsonStringPlr)
//	}
//}