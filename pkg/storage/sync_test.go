package storage

import (
	"testing"
)

func TestSync(t *testing.T) {
	if err := Sync(&mockReaderStorage{}, &mockWriterStorage{}); err != nil {
		t.Errorf("expected no error but got err: %s", err)
	}
}

func TestSyncWithReaderError(t *testing.T) {
	if err := Sync(&mockErrorGetAllStorageReaderWriter{}, nil); err == nil {
		t.Errorf("expected error but got nil")
	}
}

func TestSyncWithWriterError(t *testing.T) {
	if err := Sync(&mockReaderStorage{}, &mockErrorWriteStorageReaderWriter{}); err == nil {
		t.Errorf("expected error but got nil")
	}
}
