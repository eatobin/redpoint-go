package giftHistoryPkg

import (
	"testing"
)

var giftHistory1 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory2 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory3 = GiftHistoryTA{{Givee: "NotEven", Giver: "Close"}}
var giftHistory4 = GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
var giftHistory5 = GiftHistoryTA{{Givee: "me", Giver: "you"}}
var jsonString = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var badJsonString = "[{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}]"
var badJsonString2 = "[{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var giftPair1 = giftPair.GiftPairStruct{Givee: "me", Giver: "you"}

func TestGiftHistoryAssertEqual(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b GiftHistoryTA
		want bool
	}
	testCases := []testCase{
		{a: giftHistory1, b: giftHistory2, want: true},
		{a: giftHistory1, b: giftHistory3, want: false},
		{a: giftHistory1, b: giftHistory4, want: false},
	}
	for _, tc := range testCases {
		got := GiftHistoryAssertEqual(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("GiftHistoryAssertEqual(%v, %v): want %t, got %t",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestGiftHistoryJsonStringToGiftHistory(t *testing.T) {
	t.Parallel()
	got, err := GiftHistoryJsonStringToGiftHistory(jsonString)
	want := giftHistory1
	if err != nil {
		t.Fatalf("want no error for valid input, got: %v", err)
	}
	if !GiftHistoryAssertEqual(got, want) {
		t.Errorf("GiftHistoryJsonStringToGiftHistory(%s): want %v, got %v",
			giftHistory1, want, got)
	}
}

func TestGiftHistoryJsonStringToGiftHistoryInvalid(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    giftPair.JsonStringTA
		want GiftHistoryTA
	}
	testCases := []testCase{
		{a: badJsonString},
		{a: badJsonString2},
	}
	for _, tc := range testCases {
		_, err := GiftHistoryJsonStringToGiftHistory(tc.a)
		if err == nil {
			t.Error("want error for invalid input, got nil")
		}
	}
}

func TestGiftHistoryAddYear(t *testing.T) {
	t.Parallel()
	got := GiftHistoryAddYear("NewBee", giftHistory1)
	want := giftHistory4
	if !GiftHistoryAssertEqual(got, want) {
		t.Errorf("GiftHistoryAddYear(%s, %v): want %v, got %v", "NewBee", giftHistory1, want, got)
	}
}

func TestGiftHistoryUpdateGiftHistory(t *testing.T) {
	t.Parallel()
	got := GiftHistoryUpdateGiftHistory(0, giftPair1, giftHistory1)
	want := giftHistory5
	if !GiftHistoryAssertEqual(got, want) {
		t.Errorf("GiftHistoryUpdateGiftHistory(%d, %v, %v): want %v, got %v", 0, giftPair1, giftHistory1, want, got)
	}
}
