package storage

import "time"

type FeatureToggles map[string]FeatureToggle

type FeatureToggle struct {
	Description string    `json:"description,omitempty" yaml:"description"`
	Value       bool      `json:"value" yaml:"value"`
	ExpireTime  time.Time `json:"expireTime,omitempty" yaml:"expireTime"`
}
