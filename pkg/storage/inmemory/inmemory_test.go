package inmemory

import (
	"reflect"
	"testing"

	"github.com/trendyol/teg/pkg/storage"
)

func TestInMemorystorage(t *testing.T) {
	im := New()

	want := storage.FeatureToggle{
		Value: false,
	}

	if err := im.Set("name", want); err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}

	got, err := im.Get("name")
	if err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}

	if !reflect.DeepEqual(&want, got) {
		t.Errorf("expected %v but got %v", want, got)
	}

	if _, err := im.Get("not-found"); err == nil {
		t.Errorf("expected error but got nil")
	}

	if err := im.Set("name", want); err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}

	featureToggless, err := im.GetAll()
	if err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}

	if len(featureToggless) != 1 {
		t.Errorf("expected feature toggles len %d but got %d", 1, len(featureToggless))
	}
}

func TestInMemorystorageWithOptions(t *testing.T) {
	defaultFeatureToggles := storage.FeatureToggles{
		"name": storage.FeatureToggle{
			Value: true,
		},
	}

	im := New(WithDefaultToggles(
		defaultFeatureToggles,
	))

	featureToggless, err := im.GetAll()
	if err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}

	if len(featureToggless) != 1 {
		t.Errorf("expected feature toggles len %d but got %d", 1, len(featureToggless))
	}

}
