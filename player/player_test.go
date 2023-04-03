package player

import (
	"github.com/eatobin/redpoint-go/giftHistory"
	"github.com/eatobin/redpoint-go/giftPair"
	"reflect"
	"testing"
)

var jsonString = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var player1 = Player{PlayerName: "Paul McCartney", GiftHistory: giftHistory.History{{Givee: "GeoHar", Giver: "JohLen"}}}
var badJsonString = "{\"playerName\"\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var badJsonString2 = "{\"playerNameX\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"

//var player2 = Player{PlayerName: "Paul McCartney", GiftHistory: giftHistory.History{{Givee: "GeoHar", Giver: "JohLen"}}}
//var player3 = Player{PlayerName: "Paul McCartney", GiftHistory: giftHistory.History{{Givee: "Nope", Giver: "JohLen"}}}
//var player4 = Player{PlayerName: "Nope", GiftHistory: giftHistory.History{{Givee: "GeoHar", Giver: "JohLen"}}}
//var player5 = Player{PlayerName: "Paul McCartney", GiftHistory: giftHistory.History{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}}

func TestJsonStringToPlayer(t *testing.T) {
	type args struct {
		jsonString giftPair.JsonString
	}
	tests := []struct {
		name    string
		args    args
		want    Player
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonString}, want: player1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonStringToPlayer(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonStringToPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonStringToPlayer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
