package git

import (
	"errors"
	"testing"
)

func TestGitStorage(t *testing.T) {
	const file = `HelloEnable:
  value: true
  description: "Hello"`

	gitStorage := New("", "")
	gitStorage.GetFileContent = func(url, filePath, userName, password string) (string, error) {
		return file, nil
	}

	ft, err := gitStorage.Get("HelloEnable")
	if err != nil {
		t.Fatalf("expected no error but got err: %s", err)
	}

	if !ft.Value {
		t.Errorf("expected %v but got %v", true, ft.Value)
	}

	if ft.Description != "Hello" {
		t.Errorf("expected %v but got %v", "Hello", ft.Description)
	}

	if _, err := gitStorage.Get("NotFound"); err == nil {
		t.Errorf("expected error but got nil")
	}

	gitStorage.GetFileContent = func(url, filePath, userName, password string) (string, error) {
		return "", errors.New("system error")
	}

	if _, err := gitStorage.Get("HelloEnable"); err == nil {
		t.Errorf("expected error but got nil")
	}

	gitStorage.GetFileContent = func(url, filePath, userName, password string) (string, error) {
		return "asdf", nil
	}

	if _, err := gitStorage.Get("HelloEnable"); err == nil {
		t.Errorf("expected error but got nil")
	}
}

func TestGitWithOptions(t *testing.T) {
	gitStorage := New("", "", WithUserNameAndPassword("userName", "password"), WithAccessToken("password"))

	if gitStorage.UserName != "userName" {
		t.Errorf("expected %s but got %s", "userName", gitStorage.UserName)
	}

	if gitStorage.Password != "password" {
		t.Errorf("expected %s but got %s", "password", gitStorage.Password)
	}
}
