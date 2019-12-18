package filter

import (
	"testing"
	"time"

	"github.com/trendyol/teg/pkg/storage"
)

func TestFilterToggleDisable(t *testing.T) {
	var testCases = []struct {
		name string
		in   storage.FeatureToggle
		want bool
	}{
		{
			"Active Feature Toggle",
			storage.FeatureToggle{
				Value: true,
			},
			false,
		},
		{
			"Disable Feature Toggle",
			storage.FeatureToggle{
				Value: false,
			},
			true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToggleDisabled(tt.in); got != tt.want {
				t.Errorf("FilterToggleDisable() want = %v, got %v", tt.want, got)
				return
			}
		})
	}
}

func TestFilterToggleExpireTime(t *testing.T) {
	var testCases = []struct {
		name string
		in   storage.FeatureToggle
		want bool
	}{
		{
			"Correct Feature Toggle",
			storage.FeatureToggle{
				Value:      true,
				ExpireTime: time.Now().Add(1 * time.Minute),
			},
			false,
		},
		{
			"Correct Feature Toggle Empty Expire Time",
			storage.FeatureToggle{
				Value: true,
			},
			false,
		},
		{
			"Expried Feature Toggle",
			storage.FeatureToggle{
				Value:      true,
				ExpireTime: time.Now().Add(-1 * time.Minute),
			},
			true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToggleTimeExpired(tt.in); got != tt.want {
				t.Errorf("ToggleExpireTime() want = %v, got %v", tt.want, got)
				return
			}
		})
	}
}
