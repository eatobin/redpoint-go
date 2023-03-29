package giftHistoryPkg

import "testing"

var giftHistory1 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory2 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory3 = GiftHistoryTA{{Givee: "NotEven", Giver: "Close"}}
var giftHistory4 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
var jsonString = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"

func TestGiftHistoryAssertEqual(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b GiftHistoryTA
		want bool
	}
	testCases := []testCase{
		{a: giftHistory1, b: giftHistory2, want: true},
		//{a: giftPair1, b: giftPair3, want: false},
	}
	for _, tc := range testCases {
		got := GiftHistoryAssertEqual(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("GiftHistoryAssertEqual(%v, %v): want %t, got %t",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestGiftHistoryAssertEqualTrue(t *testing.T) {
	t.Parallel()
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory2)
	if got != true {
		t.Errorf("Got: %v Want: %v", got, true)
	}
}

func TestGiftPairAssertEqualFalse(t *testing.T) {
	t.Parallel()
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory3)
	if got != false {
		t.Errorf("Got: %v Want: %v", got, false)
	}
}

func TestGiftHistoryAssertEqualFalseExtended(t *testing.T) {
	t.Parallel()
	got := GiftHistoryAssertEqual(giftHistory1, giftHistory4)
	if got != false {
		t.Errorf("Got: %v Want: %v", got, false)
	}
}

func TestGiftHistoryJsonStringToGiftHistory(t *testing.T) {
	t.Parallel()
	got, _ := GiftHistoryJsonStringToGiftHistory(jsonString)
	want := giftHistory1
	if !GiftHistoryAssertEqual(got, want) {
		t.Errorf("GiftHistoryJsonStringToGiftHistory(%s): want %v, got %v", jsonString, want, got)
	}
}

func TestGiftHistoryAddYear(t *testing.T) {
	t.Parallel()
	got := GiftHistoryAddYear("NewBee", giftHistory1)
	want := giftHistory4
	if !GiftHistoryAssertEqual(got, want) {
		t.Errorf("GiftHistoryAddYear(%s,%v): want %v, got %v", "NewBee", giftHistory1, want, got)
	}
}

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
