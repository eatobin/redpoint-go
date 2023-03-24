package giftHistoryPkg

import "testing"

var giftHistory1 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory2 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory3 = GiftHistoryTA{{Givee: "NotEven", Giver: "Close"}}
var giftHistory4 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
var jsonString = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"

func TestGiftHistoryAssertEqualTrue(t *testing.T) {
	t.Parallel()
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory2)
	if got != true {
		t.Fatalf("Got: %v Want: %v", got, true)
	}
}

func TestGiftPairAssertEqualFalse(t *testing.T) {
	t.Parallel()
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory3)
	if got != false {
		t.Fatalf("Got: %v Want: %v", got, false)
	}
}

func TestGiftHistoryAssertEqualFalseExtended(t *testing.T) {
	t.Parallel()
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory4)
	if got != false {
		t.Fatalf("Got: %v Want: %v", got, false)
	}
}

func TestGiftHistoryJsonStringToGiftHistory(t *testing.T) {
	t.Parallel()
	got, _ := GiftHistoryJsonStringToGiftHistory(jsonString)
	want := giftHistory1
	if !GiftHistoryAssertEqual(got, want) {
		t.Fatalf("Got: %v Want: %v", got, want)
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

//	if !GiftHistoryAssertEqual(giftHistory, gotGH) {
//		t.Fatalf("JsonStringToGiftHistory(%s) == %v,\nwant %v", jsonStringGH, gotGH, giftHistory)
//	}
//}
//
//func TestGiftHistoryToJsonString(t *testing.T) {
//	jsonStringGH := "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
//	giftHistory := GiftHistory{{Givee: "GeoHar", Giver: "JohLen"}}
//	gotGHstring, _ := GiftHistoryJsonStringToGiftHistory(giftHistory)
//
//	if jsonStringGH != gotGHstring {
//		t.Fatalf("GiftHistoryJsonStringToGiftHistory(%v) == %s,\nwant %s", giftHistory, gotGHstring, jsonStringGH)
//	}
//}
