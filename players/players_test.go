package players

import (
	"testing"
)

//var jsonStringPlayers = "{\"PauMcc\":{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
//var geoHarGivee = Player{PlayerName: "George Harrison", GiftHistory: GiftHistory{{Givee: "you", Giver: "PauMcc"}}}
//var geoHarGiver = Player{PlayerName: "George Harrison", GiftHistory: GiftHistory{{Givee: "RinSta", Giver: "you"}}}
//var playersGivee = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHarGivee, "PauMcc": pauMcc}
//var playersGiver = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHarGiver, "PauMcc": pauMcc}
//var jsonStringPlr = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"

func TestComparePlayers(t *testing.T) {
	var rinSta = Player{PlayerName: "Ringo Starr", GiftHistory: GiftHistory{{Givee: "JohLen", Giver: "GeoHar"}}}
	var johLen = Player{PlayerName: "John Lennon", GiftHistory: GiftHistory{{Givee: "PauMcc", Giver: "RinSta"}}}
	var geoHar = Player{PlayerName: "George Harrison", GiftHistory: GiftHistory{{Givee: "RinSta", Giver: "PauMcc"}}}
	var pauMcc = Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}}
	var playersA = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}
	var playersB = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}
	var playersC = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar}
	var newBee = Player{PlayerName: "New Bee", GiftHistory: GiftHistory{{Givee: "NewBee", Giver: "NewBee"}}}
	var newBeePlayers = Players{"RinSta": newBee, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}

	if !ComparePlayers(playersA, playersB) {
		t.Fatalf("ComparePlayers(%v,\n%v) == %t, want %t", playersA, playersB, false, true)
	}
	if ComparePlayers(playersB, newBeePlayers) {
		t.Fatalf("ComparePlayers(%v,\n%v) == %t, want %t", playersB, newBeePlayers, true, false)
	}
	if ComparePlayers(playersB, playersC) {
		t.Fatalf("ComparePlayers(%v,\n%v) == %t, want %t", playersB, playersC, true, false)
	}
}

func TestUpdatePlayer(t *testing.T) {
	goodBee := Player{PlayerName: "New Bee", GiftHistory: GiftHistory{{Givee: "NewBee", Giver: "NewBee"}}}
	var rinSta = Player{PlayerName: "Ringo Starr", GiftHistory: GiftHistory{{Givee: "JohLen", Giver: "GeoHar"}}}
	var johLen = Player{PlayerName: "John Lennon", GiftHistory: GiftHistory{{Givee: "PauMcc", Giver: "RinSta"}}}
	var geoHar = Player{PlayerName: "George Harrison", GiftHistory: GiftHistory{{Givee: "RinSta", Giver: "PauMcc"}}}
	var pauMcc = Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}}
	var playersA = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}
	var newBee = Player{PlayerName: "New Bee", GiftHistory: GiftHistory{{Givee: "NewBee", Giver: "NewBee"}}}
	var newBeePlayers = Players{"RinSta": newBee, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}

	gotPlayers := UpdatePlayer("RinSta", goodBee, playersA)
	if !ComparePlayers(gotPlayers, newBeePlayers) {
		t.Fatalf("UpdatePlayer(%s, %v, %v) ==\n%v,\nwant %v", "RinSta", newBee, playersA, gotPlayers, newBeePlayers)
	}
}

func TestGetPlayerName(t *testing.T) {
	var rinSta = Player{PlayerName: "Ringo Starr", GiftHistory: GiftHistory{{Givee: "JohLen", Giver: "GeoHar"}}}
	var johLen = Player{PlayerName: "John Lennon", GiftHistory: GiftHistory{{Givee: "PauMcc", Giver: "RinSta"}}}
	var geoHar = Player{PlayerName: "George Harrison", GiftHistory: GiftHistory{{Givee: "RinSta", Giver: "PauMcc"}}}
	var pauMcc = Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}}
	var playersA = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}

	name := GetPlayerName("PauMcc", playersA)

	if name != "Paul McCartney" {
		t.Fatalf("GetPlayerName(%s, %v) = \n%s\nwant %s", "PaulMcc", playersA, name, "Paul McCartney")
	}
}

func TestAddYear(t *testing.T) {
	var rinSta = Player{PlayerName: "Ringo Starr", GiftHistory: GiftHistory{{Givee: "JohLen", Giver: "GeoHar"}}}
	var johLen = Player{PlayerName: "John Lennon", GiftHistory: GiftHistory{{Givee: "PauMcc", Giver: "RinSta"}}}
	var geoHar = Player{PlayerName: "George Harrison", GiftHistory: GiftHistory{{Givee: "RinSta", Giver: "PauMcc"}}}
	var pauMcc = Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}}
	var playersA = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}

	var rinStaExt = Player{PlayerName: "Ringo Starr", GiftHistory: GiftHistory{{Givee: "JohLen", Giver: "GeoHar"}, {Givee: "RinSta", Giver: "RinSta"}}}
	var johLenExt = Player{PlayerName: "John Lennon", GiftHistory: GiftHistory{{Givee: "PauMcc", Giver: "RinSta"}, {Givee: "JohLen", Giver: "JohLen"}}}
	var geoHarExt = Player{PlayerName: "George Harrison", GiftHistory: GiftHistory{{Givee: "RinSta", Giver: "PauMcc"}, {Givee: "GeoHar", Giver: "GeoHar"}}}
	var pauMccExt = Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "PauMcc", Giver: "PauMcc"}}}
	var playersExt = Players{"RinSta": rinStaExt, "JohLen": johLenExt, "GeoHar": geoHarExt, "PauMcc": pauMccExt}

	newYear := AddYear(playersA)
	if !ComparePlayers(newYear, playersExt) {
		t.Fatalf("AddYear(%v) ==\n%v,\nwant %v", playersA, newYear, playersExt)
	}
}

//it should "add a new year" in {
//assert(Players.addYear(players) == playersExt)
//}
//
//it should "return a givee and a giver" in {
//assert(Players.getGivee("GeoHar")(0)(players) == "RinSta")
//assert(Players.getGiver("GeoHar")(0)(players) == "PauMcc")
//}
//
//it should "update a givee and a giver" in {
//assert(Players.updateGivee("GeoHar")(0)("you")(players) == playersGivee)
//assert(Players.updateGiver("GeoHar")(0)("you")(players) == playersGiver)
//}
//
//it should "convert from JSON" in {
//val plrsJson: Either[Error, Map[String, Player]] = Players.jsonStringToPlayers(jsonStringPlrs)
//assert(plrsJson == Right(players))
//}
