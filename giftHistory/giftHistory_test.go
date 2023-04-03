package giftHistory

import (
	"github.com/eatobin/redpoint-go/giftPair"
	"reflect"
	"testing"
)

var giftHistory1 = HistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory2 = HistoryTA{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory3 = HistoryTA{{Givee: "NotEven", Giver: "Close"}}
var giftHistory4 = HistoryTA{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
var giftHistory5 = HistoryTA{{Givee: "me", Giver: "you"}}
var jsonString = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var badJsonString = "[{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}]"
var badJsonString2 = "[{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var giftPair1 = giftPair.Struct{Givee: "me", Giver: "you"}

//func TestGiftHistoryAssertEqual(t *testing.T) {
//	t.Parallel()
//	type testCase struct {
//		a, b HistoryTA
//		want bool
//	}
//	testCases := []testCase{
//		{a: giftHistory1, b: giftHistory2, want: true},
//		{a: giftHistory1, b: giftHistory3, want: false},
//		{a: giftHistory1, b: giftHistory4, want: false},
//	}
//	for _, tc := range testCases {
//		got := HistoryAssertEqual(tc.a, tc.b)
//		if tc.want != got {
//			t.Errorf("HistoryAssertEqual(%v, %v): want %t, got %t",
//				tc.a, tc.b, tc.want, got)
//		}
//	}
//}
//
//func TestGiftHistoryJsonStringToGiftHistory(t *testing.T) {
//	t.Parallel()
//	got, err := JsonStringToGiftHistory(jsonString)
//	want := giftHistory1
//	if err != nil {
//		t.Fatalf("want no error for valid input, got: %v", err)
//	}
//	if !HistoryAssertEqual(got, want) {
//		t.Errorf("JsonStringToGiftHistory(%s): want %v, got %v",
//			giftHistory1, want, got)
//	}
//}
//
//func TestGiftHistoryJsonStringToGiftHistoryInvalid(t *testing.T) {
//	t.Parallel()
//	type testCase struct {
//		a    giftPair.JsonStringTA
//		want HistoryTA
//	}
//	testCases := []testCase{
//		{a: badJsonString},
//		{a: badJsonString2},
//	}
//	for _, tc := range testCases {
//		_, err := JsonStringToGiftHistory(tc.a)
//		if err == nil {
//			t.Error("want error for invalid input, got nil")
//		}
//	}
//}
//
//func TestGiftHistoryAddYear(t *testing.T) {
//	t.Parallel()
//	got := AddYear("NewBee", giftHistory1)
//	want := giftHistory4
//	if !HistoryAssertEqual(got, want) {
//		t.Errorf("AddYear(%s, %v): want %v, got %v", "NewBee", giftHistory1, want, got)
//	}
//}
//
//func TestGiftHistoryUpdateGiftHistory(t *testing.T) {
//	t.Parallel()
//	got := UpdateGiftHistory(0, giftPair1, giftHistory1)
//	want := giftHistory5
//	if !HistoryAssertEqual(got, want) {
//		t.Errorf("UpdateGiftHistory(%d, %v, %v): want %v, got %v", 0, giftPair1, giftHistory1, want, got)
//	}
//}

func TestAddYear(t *testing.T) {
	t.Parallel()
	type args struct {
		playerKey   giftPair.PlayerKeyTA
		giftHistory HistoryTA
	}
	tests := []struct {
		name string
		args args
		want HistoryTA
	}{
		{
			name: "ValidInput",
			args: args{playerKey: "NewBee", giftHistory: giftHistory1},
			want: giftHistory4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddYear(tt.args.playerKey, tt.args.giftHistory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistoryAssertEqual(t *testing.T) {
	type args struct {
		a HistoryTA
		b HistoryTA
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HistoryAssertEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("HistoryAssertEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonStringToGiftHistory(t *testing.T) {
	type args struct {
		jsonString giftPair.JsonStringTA
	}
	tests := []struct {
		name    string
		args    args
		want    HistoryTA
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonStringToGiftHistory(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonStringToGiftHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonStringToGiftHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateGiftHistory(t *testing.T) {
	type args struct {
		giftYear    GiftYearTA
		giftPair    giftPair.Struct
		giftHistory HistoryTA
	}
	tests := []struct {
		name string
		args args
		want HistoryTA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateGiftHistory(tt.args.giftYear, tt.args.giftPair, tt.args.giftHistory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateGiftHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}
