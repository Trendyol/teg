package env

import "testing"

import "os"

import "github.com/trendyol/teg/pkg/storage"

func TestEnvstorage(t *testing.T) {
	env := New()

	env.Set("Name", storage.FeatureToggle{
		Value: true,
	})
	defer os.Unsetenv("TEG_NAME")

	os.Setenv("TEG_NAME_WRONG", "asdf")
	defer os.Unsetenv("TEG_NAME_WRONG")

	if _, err := env.Get("Name"); err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}

	if _, err := env.Get("NotFound"); err == nil {
		t.Errorf("expected error but got nil")
	}

	if _, err := env.Get("NameWrong"); err == nil {
		t.Errorf("expected error but got nil")
	}

	ft, err := env.GetAll()
	if err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}

	if len(ft) != 1 {
		t.Errorf("expected len %d but got %d", 1, len(ft))
	}
}

func TestEnvWithOptions(t *testing.T) {
	env := New(WithPrefix("TEST"))

	env.Set("Name", storage.FeatureToggle{
		Value: true,
	})
	defer os.Unsetenv("TEG_NAME")

	if _, err := env.Get("Name"); err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}
}
