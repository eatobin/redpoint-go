package player

import (
	"github.com/eatobin/redpoint-go/giftHistory"
	"github.com/eatobin/redpoint-go/giftPair"
	"reflect"
	"testing"
)

var jsonString = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var player1 = Struct{PlayerName: "Paul McCartney", GiftHistory: giftHistory.History{{Givee: "GeoHar", Giver: "JohLen"}}}
var badJsonString = "{\"playerName\"\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var badJsonString2 = "{\"playerNameX\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var badJsonString3 = "{\"playerName\":\"Paul McCartney\",\"giftHistoryX\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]}"
var badJsonString4 = "{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}]}"

//var player2 = Struct{PlayerName: "Paul McCartney", GiftHistory: giftHistory.History{{Givee: "GeoHar", Giver: "JohLen"}}}
//var player3 = Struct{PlayerName: "Paul McCartney", GiftHistory: giftHistory.History{{Givee: "Nope", Giver: "JohLen"}}}
//var player4 = Struct{PlayerName: "Nope", GiftHistory: giftHistory.History{{Givee: "GeoHar", Giver: "JohLen"}}}
//var player5 = Struct{PlayerName: "Paul McCartney", GiftHistory: giftHistory.History{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}}

func TestJsonStringToPlayer(t *testing.T) {
	t.Parallel()
	type args struct {
		jsonString giftPair.JsonString
	}
	tests := []struct {
		name    string
		args    args
		want    Struct
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonString}, want: player1, wantErr: false},
		{name: "InvalidInputBadJSON", args: args{badJsonString}, want: Struct{}, wantErr: true},
		{name: "InvalidInputBadPlayerName", args: args{badJsonString2}, want: Struct{}, wantErr: true},
		{name: "InvalidInputBadHistoryName", args: args{badJsonString3}, want: Struct{}, wantErr: true},
		{name: "InvalidInputBadHistoryField", args: args{badJsonString4}, want: Struct{}, wantErr: true},
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
