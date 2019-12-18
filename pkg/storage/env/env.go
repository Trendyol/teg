package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/trendyol/teg/pkg/storage"
)

const defaultPrefix = "TEG"

type osEnvironmentStorage struct {
	Prefix string
}

var _ storage.ReaderWriter = (*osEnvironmentStorage)(nil)

func New(setters ...option) *osEnvironmentStorage {
	args := &options{
		Prefix: defaultPrefix,
	}

	for _, setter := range setters {
		setter(args)
	}

	return &osEnvironmentStorage{
		Prefix: args.Prefix,
	}
}

func (i *osEnvironmentStorage) Get(name string) (*storage.FeatureToggle, error) {
	envName := fmt.Sprintf("%s_%s", i.Prefix, strcase.ToScreamingSnake(name))

	envValue, ok := os.LookupEnv(envName)
	if !ok {
		return nil, fmt.Errorf("couldn't found (%s) feature toggle in env", name)
	}

	val, err := strconv.ParseBool(envValue)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse (%s) value to bool. err: %w", envValue, err)
	}

	return &storage.FeatureToggle{
		Value: val,
	}, nil
}

func (i *osEnvironmentStorage) GetAll() (storage.FeatureToggles, error) {
	toggles := make(storage.FeatureToggles)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) != 2 {
			continue
		}

		// Env name with prefix. for example: TEG_BLA_BLA
		rawEnvName := pair[0]
		if !strings.HasPrefix(rawEnvName, "TEG") {
			continue
		}

		// Remove prefix. for example: BLA_BLA
		rawEnvName = rawEnvName[len(i.Prefix):]

		name := strcase.ToCamel(strings.ToLower(rawEnvName))
		val, err := strconv.ParseBool(pair[1])
		if err != nil {
			continue
		}

		toggles[name] = storage.FeatureToggle{
			Value: val,
		}
	}

	return toggles, nil
}

func (i *osEnvironmentStorage) Set(name string, toggle storage.FeatureToggle) error {
	envName := fmt.Sprintf("%s_%s", i.Prefix, strcase.ToScreamingSnake(name))
	envValue := strconv.FormatBool(toggle.Value)

	return os.Setenv(envName, envValue)
}
