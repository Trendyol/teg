package storage

import (
	"errors"
)

type mockReaderStorage struct{}

func (e *mockReaderStorage) Get(_ string) (*FeatureToggle, error) { return nil, nil }
func (e *mockReaderStorage) GetAll() (FeatureToggles, error) {
	return FeatureToggles{
		"Name": FeatureToggle{
			Value: true,
		},
	}, nil
}

type mockWriterStorage struct{}

func (e *mockWriterStorage) Set(_ string, _ FeatureToggle) error { return nil }

type mockErrorGetAllStorageReaderWriter struct{}

func (e *mockErrorGetAllStorageReaderWriter) Get(_ string) (*FeatureToggle, error) { return nil, nil }
func (e *mockErrorGetAllStorageReaderWriter) GetAll() (FeatureToggles, error) {
	return nil, errors.New("system error")
}

type mockErrorWriteStorageReaderWriter struct{}

func (e *mockErrorWriteStorageReaderWriter) Get(_ string) (*FeatureToggle, error) { return nil, nil }
func (e *mockErrorWriteStorageReaderWriter) GetAll() (FeatureToggles, error)      { return nil, nil }
func (e *mockErrorWriteStorageReaderWriter) Set(_ string, _ FeatureToggle) error {
	return errors.New("system error")
}
