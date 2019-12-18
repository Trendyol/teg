package storage

import "fmt"

func Sync(reader Reader, writer Writer) error {
	toggles, err := reader.GetAll()
	if err != nil {
		return fmt.Errorf("couldn't read all of feature toggles. err: %w", err)
	}

	for name, toggle := range toggles {
		if err := writer.Set(name, toggle); err != nil {
			return fmt.Errorf("couldn't set (%s) feature toggle. err: %w", name, err)
		}
	}

	return nil
}
