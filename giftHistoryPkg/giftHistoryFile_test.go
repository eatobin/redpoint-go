package giftHistoryPkg

import (
	"github.com/eatobin/redpoint-go/giftPairPkg"
	"testing"
)

func TestGiftHistoryAssertEqual(t *testing.T) {
	t.Parallel()
	giftHistory1 := GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory2 := GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory3 := GiftHistoryTA{{Givee: "NotEven", Giver: "Close"}}
	giftHistory4 := GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
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
	jsonString := "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
	giftHistory1 := GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
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
	badJsonString := "[{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}]"
	badJsonString2 := "[{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}]"
	type testCase struct {
		a    giftPairPkg.JsonStringTA
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
	giftHistory1 := GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory4 := GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
	got := GiftHistoryAddYear("NewBee", giftHistory1)
	want := giftHistory4
	if !GiftHistoryAssertEqual(got, want) {
		t.Errorf("GiftHistoryAddYear(%s, %v): want %v, got %v", "NewBee", giftHistory1, want, got)
	}
}

func TestGiftHistoryUpdateGiftHistory(t *testing.T) {
	t.Parallel()
	giftHistory1 := GiftHistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
	giftHistory5 := GiftHistoryTA{{Givee: "me", Giver: "you"}}
	giftPair1 := giftPairPkg.GiftPairStruct{Givee: "me", Giver: "you"}
	got := GiftHistoryUpdateGiftHistory(0, giftPair1, giftHistory1)
	want := giftHistory5
	if !GiftHistoryAssertEqual(got, want) {
		t.Errorf("GiftHistoryUpdateGiftHistory(%d, %v, %v): want %v, got %v", 0, giftPair1, giftHistory1, want, got)
	}
}
