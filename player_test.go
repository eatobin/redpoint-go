package main

import (
	"reflect"
	"testing"
)

var jsonStringPlr = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var playerPlr1 = PlayerStruct{PlayerName: "Paul McCartney", GiftHistory: History{{Givee: "GeoHar", Giver: "JohLen"}}}
var badJsonStringPlr = "{\"playerName\"\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var badJsonStringPlr2 = "{\"playerNameX\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var badJsonStringPlr3 = "{\"playerName\":\"Paul McCartney\",\"giftHistoryX\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var badJsonStringPlr4 = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var playerPlr2 = PlayerStruct{PlayerName: "Paul McCartney", GiftHistory: History{{Givee: "nope", Giver: "yup"}}}

func TestPlayerJsonStringToPlayer(t *testing.T) {
	t.Parallel()
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerStruct
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonStringPlr}, want: playerPlr1, wantErr: false},
		{name: "InvalidInputBadJSON", args: args{badJsonStringPlr}, want: PlayerStruct{}, wantErr: true},
		{name: "InvalidInputBadPlayerName", args: args{badJsonStringPlr2}, want: PlayerStruct{}, wantErr: true},
		{name: "InvalidInputBadHistoryName", args: args{badJsonStringPlr3}, want: PlayerStruct{}, wantErr: true},
		{name: "InvalidInputBadHistoryField", args: args{badJsonStringPlr4}, want: PlayerStruct{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PlayerJsonStringToPlayer(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlayerJsonStringToPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayerJsonStringToPlayer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerStruct_PlayerUpdateGiftHistory(t *testing.T) {
	type fields struct {
		PlayerName  string
		GiftHistory History
	}
	type args struct {
		giftHistory History
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PlayerStruct
	}{
		{
			name:   "ValidInput",
			fields: fields(playerPlr1),
			args:   args{History{{Givee: "nope", Giver: "yup"}}},
			want:   playerPlr2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := PlayerStruct{
				PlayerName:  tt.fields.PlayerName,
				GiftHistory: tt.fields.GiftHistory,
			}
			if got := player.PlayerUpdateGiftHistory(tt.args.giftHistory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayerUpdateGiftHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerStruct_PlayerString(t *testing.T) {
	type fields struct {
		PlayerName  string
		GiftHistory History
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "ValidInput",
			fields: fields(playerPlr1),
			want:   "{PlayerName: Paul McCartney, GiftHistory: [{GeoHar JohLen}]}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := PlayerStruct{
				PlayerName:  tt.fields.PlayerName,
				GiftHistory: tt.fields.GiftHistory,
			}
			if got := player.PlayerString(); got != tt.want {
				t.Errorf("PlayerString() = %v, want %v", got, tt.want)
			}
		})
	}
}
