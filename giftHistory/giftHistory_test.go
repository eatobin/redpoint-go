package giftHistory

import (
	"testing"
)

////var jsonStringGH = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b GiftHistory) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if *v != *b[i] {
			return false
		}
	}
	return true
}

func TestAddYear(t *testing.T) {
	giftHistory := GiftHistory{&GiftPair{Givee: "GeoHar", Giver: "JohLen"}}
	gotGHAY := AddYear("NewBee", giftHistory)
	giftHistoryExtended := GiftHistory{&GiftPair{Givee: "GeoHar", Giver: "JohLen"}, &GiftPair{Givee: "NewBee", Giver: "NewBee"}}
	if !equal(gotGHAY, giftHistoryExtended) {
		t.Fatalf("AddYear(%s, %v) == %v, want %v", "NewBee", giftHistory, gotGHAY, giftHistoryExtended)
	}
}

func TestUpdateGiftHistory(t *testing.T) {
	giftHistoryBase := GiftHistory{&GiftPair{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory := GiftHistory{&GiftPair{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistoryMeYou := GiftHistory{&GiftPair{Givee: "me", Giver: "you"}}
	UpdateGiftHistory(0, &GiftPair{Givee: "me", Giver: "you"}, giftHistory)
	if !equal(giftHistory, giftHistoryMeYou) {
		t.Fatalf("UpdateGiftHistory(%d, %v, %v) == %v, want %v", 0, GiftPair{Givee: "me", Giver: "you"}, giftHistoryBase, giftHistory, giftHistoryMeYou)
	}
}
