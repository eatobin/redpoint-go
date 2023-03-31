package playerPkg

import (
	"github.com/eatobin/redpoint-go/giftHistoryPkg"
	"testing"
)

//var jsonStringPlr = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"

var playerD = Player{PlayerName: "Nope", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}}
var playerE = Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}}

//func TestComparePlayer(t *testing.T) {
//	t.Parallel()
//	var playerA = Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}}
//	var playerB = Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}}
//	if !PlayerAssertEqual(playerA, playerB) {
//		t.Errorf("ComparePlayer(%v,\n%v) == %t, want %t", playerA, playerB, false, true)
//	}
//if ComparePlayer(playerB, playerC) {
//	t.Fatalf("ComparePlayer(%v,\n%v) == %t, want %t", playerB, playerC, true, false)
//}
//if ComparePlayer(playerB, playerD) {
//	t.Fatalf("ComparePlayer(%v,\n%v) == %t, want %t", playerB, playerD, true, false)
//}
//if ComparePlayer(playerB, playerE) {
//	t.Fatalf("ComparePlayer(%v,\n%v) == %t, want %t", playerB, playerE, true, false)
//}
//}

func TestPlayerAssertEqual(t *testing.T) {
	t.Parallel()
	playerA := Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}}
	playerB := Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}}
	playerC := Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "NotEven", Giver: "Close"}}}
	type testCase struct {
		a, b Player
		want bool
	}
	testCases := []testCase{
		{a: playerA, b: playerB, want: true},
		{a: playerA, b: playerC, want: false},
	}
	for _, tc := range testCases {
		got := PlayerAssertEqual(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("PlayerAssertEqual(%v, %v): want %t, got %t",
				tc.a, tc.b, tc.want, got)
		}
	}
}

//
//func TestUpdateGiftHistory(t *testing.T) {
//	newGH := GiftHistory{{Givee: "nope", Giver: "yup"}}
//	newPlayer := Player{PlayerName: "Paul McCartney", GiftHistory: GiftHistory{{Givee: "nope", Giver: "yup"}}}.String()
//	gotPlayer := playerA.UpdateGiftHistory(newGH).String()
//	if gotPlayer != newPlayer {
//		t.Fatalf("(%v) UpdateGiftHistory(%v) ==\n%v,\nwant %v", playerA, newGH, gotPlayer, newPlayer)
//	}
//}
//
//func TestJsonStringToPlayer(t *testing.T) {
//	gotPlayer, _ := JsonStringToPlayer(jsonStringPlr)
//	if gotPlayer.String() != playerA.String() {
//		t.Fatalf("JsonStringToGiftPlayer(%v) ==\n%v,\nwant %v", jsonStringPlr, gotPlayer.String(), playerA.String())
//	}
//}
//
//func TestPlayerToJsonString(t *testing.T) {
//	gotString, _ := playerA.PlayerToJsonString()
//	if gotString != jsonStringPlr {
//		t.Fatalf("(%v) JsonStringToPlayer() ==\n%v,\nwant %v", playerA, gotString, jsonStringPlr)
//	}
//}
