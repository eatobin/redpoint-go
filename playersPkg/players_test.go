package playersPkg

import (
	"github.com/eatobin/redpoint-go/giftHistoryPkg"
	"github.com/eatobin/redpoint-go/playerPkg"
	"reflect"
	"testing"
)

var jsonString = "{\"PauMcc\":{\"playerName\":\"Paul McCartney\",\"giftHistory\":[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]},\"GeoHar\":{\"playerName\":\"George Harrison\",\"giftHistory\":[{\"givee\":\"RinSta\",\"giver\":\"PauMcc\"}]},\"JohLen\":{\"playerName\":\"John Lennon\",\"giftHistory\":[{\"givee\":\"PauMcc\",\"giver\":\"RinSta\"}]},\"RinSta\":{\"playerName\":\"Ringo Starr\",\"giftHistory\":[{\"givee\":\"JohLen\",\"giver\":\"GeoHar\"}]}}"
var rinSta = playerPkg.StructPlayer{PlayerName: "Ringo Starr", GiftHistory: giftHistoryPkg.History{{Givee: "JohLen", Giver: "GeoHar"}}}
var johLen = playerPkg.StructPlayer{PlayerName: "John Lennon", GiftHistory: giftHistoryPkg.History{{Givee: "PauMcc", Giver: "RinSta"}}}
var geoHar = playerPkg.StructPlayer{PlayerName: "George Harrison", GiftHistory: giftHistoryPkg.History{{Givee: "RinSta", Giver: "PauMcc"}}}
var pauMcc = playerPkg.StructPlayer{PlayerName: "Paul McCartney", GiftHistory: giftHistoryPkg.History{{Givee: "GeoHar", Giver: "JohLen"}}}
var players = Players{"RinSta": rinSta, "JohLen": johLen, "GeoHar": geoHar, "PauMcc": pauMcc}

func TestJsonStringToPlayers(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    Players
		wantErr bool
	}{
		{name: "ValidInput", args: args{jsonString}, want: players, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonStringToPlayers(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonStringToPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonStringToPlayers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
