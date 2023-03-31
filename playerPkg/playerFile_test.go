package playerPkg

import (
	"github.com/eatobin/redpoint-go/giftHistoryPkg"
	"testing"
)

//var jsonStringPlr = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"

func TestPlayerAssertEqual(t *testing.T) {
	t.Parallel()
	player1 := Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}}
	player2 := Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}}
	player3 := Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "Nope", Giver: "JohLen"}}}
	player4 := Player{PlayerName: "Nope", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}}
	player5 := Player{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}}
	type testCase struct {
		a, b Player
		want bool
	}
	testCases := []testCase{
		{a: player1, b: player2, want: true},
		{a: player1, b: player3, want: false},
		{a: player1, b: player4, want: false},
		{a: player1, b: player5, want: false},
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
