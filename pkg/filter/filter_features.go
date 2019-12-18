package filter

import (
	"time"

	"github.com/trendyol/teg/pkg/storage"
)

func ToggleDisabled(t storage.FeatureToggle) bool {
	return !t.Value
}

func ToggleTimeExpired(t storage.FeatureToggle) bool {
	return !t.ExpireTime.IsZero() && time.Now().After(t.ExpireTime)
}
