package api

import (
	"reflect"
	"testing"
)

func TestGetAfricasTalkingSettings(t *testing.T) {
	type args struct {
		apiKey      string
		username    string
		URLendPoint string
	}
	tests := []struct {
		name string
		args args
		want *AtClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAfricasTalkingSettings(tt.args.apiKey, tt.args.username, tt.args.URLendPoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAfricasTalkingSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}
