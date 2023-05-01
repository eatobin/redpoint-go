package main

import (
	"reflect"
	"testing"
)

var jsonStringPlrs = "{\"PauMcc\":{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
var badJsonStringPlrs = "{\"PauMcc\":{\"playerName\"\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
var badJsonStringPlrs2 = "{\"\":{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
var badJsonStringPlrs3 = "{\"PauMcc\":{\"playerNameX\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
var badJsonStringPlrs4 = "{\"PauMcc\":{\"playerName\":\"Paul McCartney\",\"giftHistoryX\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
var badJsonStringPlrs5 = "{\"PauMcc\":{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
var rinSta = PlayerStruct{PlayerName: "Ringo Starr", GiftHistory: History{{Givee: "JohLen", Giver: "GeoHar"}}}
var johLen = PlayerStruct{PlayerName: "John Lennon", GiftHistory: History{{Givee: "PauMcc", Giver: "RinSta"}}}
var geoHar = PlayerStruct{PlayerName: "George Harrison", GiftHistory: History{{Givee: "RinSta", Giver: "PauMcc"}}}
var pauMcc = PlayerStruct{PlayerName: "Paul McCartney", GiftHistory: History{{Givee: "GeoHar", Giver: "JohLen"}}}
var newBee = PlayerStruct{PlayerName: "New Bee", GiftHistory: History{{Givee: "NewBee", Giver: "NewBee"}}}
var players = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}
var newBeePlayers = Players{"RinSta": newBee, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}

func TestPlayersJsonStringToPlayers(t *testing.T) {
	t.Parallel()
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    Players
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonStringPlrs}, want: players, wantErr: false},
		{name: "InvalidInputBadJSON", args: args{badJsonStringPlrs}, want: Players{}, wantErr: true},
		{name: "InvalidInputBadPlayerKey", args: args{badJsonStringPlrs2}, want: Players{}, wantErr: true},
		{name: "InvalidInputBadPlayerName", args: args{badJsonStringPlrs3}, want: Players{}, wantErr: true},
		{name: "InvalidInputBadHistoryName", args: args{badJsonStringPlrs4}, want: Players{}, wantErr: true},
		{name: "InvalidInputBadHistoryField", args: args{badJsonStringPlrs5}, want: Players{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PlayersJsonStringToPlayers(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlayersJsonStringToPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayersJsonStringToPlayers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playersUpdatePlayer(t *testing.T) {
	type args struct {
		playerKey string
		player    PlayerStruct
		players   Players
	}
	tests := []struct {
		name string
		args args
		want Players
	}{
		{
			name: "ValidInput",
			args: args{
				"RinSta",
				PlayerStruct{PlayerName: "New Bee", GiftHistory: History{{Givee: "NewBee", Giver: "NewBee"}}},
				players,
			},
			want: newBeePlayers},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playersUpdatePlayer(tt.args.playerKey, tt.args.player, tt.args.players); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("playersUpdatePlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
