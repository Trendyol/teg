package storage

type Reader interface {
	Get(name string) (*FeatureToggle, error)
	GetAll() (FeatureToggles, error)
}

type Writer interface {
	Set(name string, toggle FeatureToggle) error
}

type ReaderWriter interface {
	Reader
	Writer
}
