package filter

import (
	"testing"
	"time"

	"github.com/trendyol/teg/pkg/storage"
)

func TestFilter(t *testing.T) {
	var testCases = []struct {
		name  string
		in    storage.FeatureToggle
		funcs []Func
		want  bool
	}{
		{
			"Filter Disable Feature Toggle",
			storage.FeatureToggle{
				Value: false,
			},
			[]Func{ToggleDisabled},
			false,
		},
		{
			"Filter Expired Feature Toggle",
			storage.FeatureToggle{
				Value:      true,
				ExpireTime: time.Now().Add(-1 * time.Minute),
			},
			[]Func{ToggleDisabled, ToggleTimeExpired},
			false,
		},
		{
			"Empty Filter Feature Toggle",
			storage.FeatureToggle{
				Value:      true,
				ExpireTime: time.Now().Add(1 * time.Minute),
			},
			[]Func{},
			true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.in, tt.funcs...); got != tt.want {
				t.Errorf("Filter() want = %v, got %v", tt.want, got)
				return
			}
		})
	}
}
