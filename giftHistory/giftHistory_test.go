package giftHistory

import (
	"testing"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []GiftPair) bool {
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
	giftHistoryBase := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistoryExtended := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}

	if !equal(AddYear("NewBee", giftHistory), giftHistoryExtended) {
		t.Fatalf("AddYear(%s, %v) == %v,\nwant %v", "NewBee", giftHistoryBase, giftHistory, giftHistoryExtended)
	}
}

func TestUpdateGiftHistory(t *testing.T) {
	giftHistoryBase := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistoryMeYou := GiftHistory{{Givee: "me", Giver: "you"}}

	if !equal(UpdateGiftHistory(0, GiftPair{Givee: "me", Giver: "you"}, giftHistory), giftHistoryMeYou) {
		t.Fatalf("UpdateGiftHistory(%d, %v, %v) == %v,\nwant %v", 0, GiftPair{Givee: "me", Giver: "you"}, giftHistoryBase, giftHistory, giftHistoryMeYou)
	}
}

func TestJsonStringToGiftHistory(t *testing.T) {
	jsonStringGH := "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	gotGH, _ := JsonStringToGiftHistory(jsonStringGH)

	if !equal(giftHistory, gotGH) {
		t.Fatalf("JsonStringToGiftHistory(%s) == %v,\nwant %v", jsonStringGH, gotGH, giftHistory)
	}
}

func TestGiftHistoryToJsonString(t *testing.T) {
	jsonStringGH := "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	gotGHstring, _ := GHToJsonString(giftHistory)

	if jsonStringGH != gotGHstring {
		t.Fatalf("GHToJsonString(%v) == %s,\nwant %s", giftHistory, gotGHstring, jsonStringGH)
	}
}
