package filter

import "github.com/trendyol/teg/pkg/storage"

type Func func(storage.FeatureToggle) bool

func Filter(toggle storage.FeatureToggle, filterFuncs ...Func) bool {
	for _, filterFunc := range filterFuncs {
		if filterFunc(toggle) {
			return false
		}
	}

	return true
}

func RequiredFilters() []Func {
	return []Func{ToggleDisabled, ToggleTimeExpired}
}
