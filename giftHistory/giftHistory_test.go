package giftHistory

import (
	"testing"
)

//var jsonStringGH = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var giftHistory = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
var gotGHAY = AddYear("NewBee", giftHistory)
var gotGHUGH = UpdateGiftHistory(0, GiftPair{Givee: "me", Giver: "you"}, giftHistory)
var giftHistoryExtended = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
var giftHistoryMeYou = GiftHistory{{"me", "you"}}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b GiftHistory) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestAddYear(t *testing.T) {
	if !equal(gotGHAY, giftHistoryExtended) {
		t.Fatalf("AddYear(%s, %v) == %v, want %v", "NewBee", giftHistory, gotGHAY, giftHistoryExtended)
	}
}

//it should "return an updated giftHistory" in {
//assert(GiftHistory.UpdateGiftHistory(0)(GiftPair("me", "you"))(giftHistory) == Vector(GiftPair("me", "you")))
//}
func TestUpdateGiftHistory(t *testing.T) {
	if !equal(gotGHUGH, giftHistoryMeYou) {
		t.Fatalf("UpdateGiftHistory(%d, %v, %v) == %v, want %v", 0, GiftPair{Givee: "me", Giver: "you"}, giftHistory, gotGHUGH, giftHistoryMeYou)
	}
}
