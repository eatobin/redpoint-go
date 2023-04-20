package main

import (
	"reflect"
	"testing"
)

var giftHistory1GH = History{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory2GH = History{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
var giftHistory3GH = History{{Givee: "me", Giver: "you"}}
var jsonStringGH = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var badJsonStringGH = "[{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}]"
var badJsonString2GH = "[{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var giftPair1GH = GiftPairStruct{Givee: "me", Giver: "you"}

func TestGiftHistoryJsonStringToGiftHistory(t *testing.T) {
	t.Parallel()
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    History
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonStringGH}, want: giftHistory1GH, wantErr: false},
		{name: "InvalidInputBadJSON", args: args{badJsonStringGH}, want: History{}, wantErr: true},
		{name: "InvalidInputBadFieldName", args: args{badJsonString2GH}, want: History{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GiftHistoryJsonStringToGiftHistory(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GiftHistoryJsonStringToGiftHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftHistoryJsonStringToGiftHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistory_GiftHistoryAddYear(t *testing.T) {
	t.Parallel()
	type args struct {
		playerKey string
	}
	tests := []struct {
		name        string
		giftHistory History
		args        args
		want        History
	}{
		{name: "ValidInput", giftHistory: giftHistory1GH, args: args{playerKey: "NewBee"}, want: giftHistory2GH},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.giftHistory.GiftHistoryAddYear(tt.args.playerKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftHistoryAddYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistory_GiftHistoryUpdateGiftHistory(t *testing.T) {
	t.Parallel()
	type args struct {
		giftYear int
		giftPair GiftPairStruct
	}
	tests := []struct {
		name        string
		giftHistory History
		args        args
		want        History
	}{
		{
			name:        "ValidInput",
			giftHistory: giftHistory1GH,
			args: args{
				giftYear: 0,
				giftPair: giftPair1GH,
			},
			want: giftHistory3GH},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.giftHistory.GiftHistoryUpdateGiftHistory(tt.args.giftYear, tt.args.giftPair); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiftHistoryUpdateGiftHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}
