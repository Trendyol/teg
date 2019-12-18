package git

import (
	"errors"
	"fmt"

	"github.com/trendyol/teg/pkg/storage"
	"gopkg.in/yaml.v2"
)

type gitStorage struct {
	GetFileContent GetFileContentFunc
	repositoryURL  string
	filePath       string
	*options
}

var _ storage.ReaderWriter = (*gitStorage)(nil)

func New(repositoryURL, filePath string, setters ...option) *gitStorage {
	args := &options{}

	for _, setter := range setters {
		setter(args)
	}

	return &gitStorage{
		GetFileContent: GetFileContent,
		repositoryURL:  repositoryURL,
		filePath:       filePath,
		options:        args,
	}
}

func (i *gitStorage) Get(name string) (*storage.FeatureToggle, error) {
	toggles, err := i.GetAll()
	if err != nil {
		return nil, err
	}

	val, ok := toggles[name]
	if !ok {
		return nil, fmt.Errorf("couldn't found (%s) feature toggle", name)
	}

	return &val, nil
}

func (i *gitStorage) GetAll() (storage.FeatureToggles, error) {
	content, err := i.GetFileContent(i.repositoryURL, i.filePath, i.UserName, i.Password)
	if err != nil {
		return nil, err
	}

	toggles := make(storage.FeatureToggles)

	if err = yaml.Unmarshal([]byte(content), &toggles); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal yaml content. err: %w", err)
	}

	return toggles, nil
}

func (i *gitStorage) Set(name string, toggle storage.FeatureToggle) error {
	return errors.New("Not Implemented")
}
