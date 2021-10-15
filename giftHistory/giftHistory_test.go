package giftHistory

import "testing"

var ghA = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
var ghB = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
var ghC = GiftHistory{{Givee: "NotEven", Giver: "Close"}}
var giftHistoryExtended = GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}

func TestCompareGiftHistory(t *testing.T) {
	if !CompareGiftHistory(ghA, ghB) {
		t.Fatalf("CompareGiftHistory(%v, %v) == %t, want %t", ghA, ghB, false, true)
	}
	if CompareGiftHistory(ghB, ghC) {
		t.Fatalf("CompareGiftPair(%v, %v) == %t, want %t", ghB, ghC, true, false)
	}
	if CompareGiftHistory(ghA, giftHistoryExtended) {
		t.Fatalf("CompareGiftPair(%v, %v) == %t, want %t", ghA, giftHistoryExtended, true, false)
	}
}

func TestAddYear(t *testing.T) {
	gotAdd := AddYear("NewBee", ghB)
	if !CompareGiftHistory(gotAdd, giftHistoryExtended) {
		t.Fatalf("AddYear(%s, %v) == %v,\nwant %v", "NewBee", ghA, gotAdd, giftHistoryExtended)
	}
}

//func TestUpdateGiftHistory(t *testing.T) {
//	giftHistoryBase := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
//	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
//	giftHistoryMeYou := GiftHistory{{Givee: "me", Giver: "you"}}
//
//	if !GHcompare(UpdateGiftHistory(0, GiftPair{Givee: "me", Giver: "you"}, giftHistory), giftHistoryMeYou) {
//		t.Fatalf("UpdateGiftHistory(%d, %v, %v) == %v,\nwant %v", 0, GiftPair{Givee: "me", Giver: "you"}, giftHistoryBase, giftHistory, giftHistoryMeYou)
//	}
//}
//
//func TestJsonStringToGiftHistory(t *testing.T) {
//	jsonStringGH := "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
//	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
//	gotGH, _ := JsonStringToGiftHistory(jsonStringGH)
//
//	if !GHcompare(giftHistory, gotGH) {
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
