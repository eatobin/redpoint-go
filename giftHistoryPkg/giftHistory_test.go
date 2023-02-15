package giftHistoryPkg

import (
	"testing"
)

var giftHistory1 = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory2 = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory3 = GiftHistory{{Givee: "NotEven", Giver: "Close"}}
var giftHistory4 = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}

func TestGiftHistoryAssertEqualTrue(t *testing.T) {
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory2)
	if got != true {
		t.Fatalf("Got: %v Want: %v", got, true)
	}
}

func TestGiftHistoryAssertEqualTrueExtended(t *testing.T) {
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory4)
	if got != false {
		t.Fatalf("Got: %v Want: %v", got, false)
	}
}

func TestGiftPairAssertEqualFalse(t *testing.T) {
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory3)
	if got != false {
		t.Fatalf("Got: %v Want: %v", got, false)
	}
}

//func TestAddYear(t *testing.T) {
//	gotAdd := AddYear("NewBee", giftHistory2)
//	if !GiftHistoryAssertEqual(gotAdd, giftHistory4) {
//		t.Fatalf("AddYear(%s, %v) == %v,\nwant %v", "NewBee", giftHistory, gotAdd, giftHistory4)
//	}
//}
//
//func TestUpdateGiftHistory(t *testing.T) {
//	giftHistoryBase := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
//	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
//	giftHistoryMeYou := GiftHistory{{Givee: "me", Giver: "you"}}
//
//	if !GiftHistoryAssertEqual(UpdateGiftHistory(0, giftPairPkg.GiftPair{Givee: "me", Giver: "you"}, giftHistory), giftHistoryMeYou) {
//		t.Fatalf("UpdateGiftHistory(%d, %v, %v) == %v,\nwant %v", 0, giftPairPkg.GiftPair{Givee: "me", Giver: "you"}, giftHistoryBase, giftHistory, giftHistoryMeYou)
//	}
//}
//
//func TestJsonStringToGiftHistory(t *testing.T) {
//	jsonStringGH := "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
//	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
//	gotGH, _ := JsonStringToGiftHistory(jsonStringGH)
//
//	if !GiftHistoryAssertEqual(giftHistory, gotGH) {
//		t.Fatalf("JsonStringToGiftHistory(%s) == %v,\nwant %v", jsonStringGH, gotGH, giftHistory)
//	}
//}
//
//func TestGiftHistoryToJsonString(t *testing.T) {
//	jsonStringGH := "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
//	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
//	gotGHstring, _ := GHToJsonString(giftHistory)
//
//	if jsonStringGH != gotGHstring {
//		t.Fatalf("GHToJsonString(%v) == %s,\nwant %s", giftHistory, gotGHstring, jsonStringGH)
//	}
//}
