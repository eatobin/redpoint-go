package giftHistoryPkg

import (
	"github.com/eatobin/redpoint-go/giftPairPkg"
	"reflect"
	"testing"
)

var giftHistory1 = History{{Givee: "GeoHar", Giver: "JohLen"}}
var giftHistory2 = History{{Givee: "GeoHar", Giver: "JohLen"}, {Givee: "NewBee", Giver: "NewBee"}}
var giftHistory3 = History{{Givee: "me", Giver: "you"}}
var jsonString = "[{\"givee\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var badJsonString = "[{\"givee\"\"GeoHar\",\"giver\":\"JohLen\"}]"
var badJsonString2 = "[{\"giveeX\":\"GeoHar\",\"giver\":\"JohLen\"}]"
var giftPair1 = giftPairPkg.StructGiftPair{Givee: "me", Giver: "you"}

//func TestAddYear(t *testing.T) {
//	t.Parallel()
//	type args struct {
//		playerKey   giftPair.PlayerKey
//		giftHistory History
//	}
//	tests := []struct {
//		name string
//		args args
//		want History
//	}{
//		{
//			name: "ValidInput",
//			args: args{playerKey: "NewBee", giftHistory: giftHistory1},
//			want: giftHistory2,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := AddYear(tt.args.playerKey, tt.args.giftHistory); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("AddYear() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestJsonStringToGiftHistory(t *testing.T) {
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
		{name: "ValidInput", args: args{jsonString}, want: giftHistory1, wantErr: false},
		{name: "InvalidInputBadJSON", args: args{badJsonString}, want: History{}, wantErr: true},
		{name: "InvalidInputBadFieldName", args: args{badJsonString2}, want: History{}, wantErr: true},
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

//func TestUpdateGiftHistory(t *testing.T) {
//	t.Parallel()
//	type args struct {
//		giftYear    GiftYear
//		giftPair    giftPair.StructGiftPair
//		giftHistory History
//	}
//	tests := []struct {
//		name string
//		args args
//		want History
//	}{
//		{
//			name: "ValidInput",
//			args: args{
//				giftYear:    0,
//				giftPair:    giftPair1,
//				giftHistory: giftHistory1,
//			},
//			want: giftHistory3,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := UpdateGiftHistory(tt.args.giftYear, tt.args.giftPair, tt.args.giftHistory); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("UpdateGiftHistory() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
