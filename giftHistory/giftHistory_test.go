package giftHistory

import (
	"testing"
)

func TestAddYear(t *testing.T) {
	giftHistoryBase := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistoryExtended := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}

	if !GHcompare(AddYear("NewBee", giftHistory), giftHistoryExtended) {
		t.Fatalf("AddYear(%s, %v) == %v,\nwant %v", "NewBee", giftHistoryBase, giftHistory, giftHistoryExtended)
	}
}

func TestUpdateGiftHistory(t *testing.T) {
	giftHistoryBase := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistoryMeYou := GiftHistory{{Givee: "me", Giver: "you"}}

	if !GHcompare(UpdateGiftHistory(0, GiftPair{Givee: "me", Giver: "you"}, giftHistory), giftHistoryMeYou) {
		t.Fatalf("UpdateGiftHistory(%d, %v, %v) == %v,\nwant %v", 0, GiftPair{Givee: "me", Giver: "you"}, giftHistoryBase, giftHistory, giftHistoryMeYou)
	}
}

func TestJsonStringToGiftHistory(t *testing.T) {
	jsonStringGH := "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
	gotGH, _ := JsonStringToGiftHistory(jsonStringGH)

	if !GHcompare(giftHistory, gotGH) {
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
