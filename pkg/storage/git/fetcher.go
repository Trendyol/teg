package git

import (
	"fmt"

	gitFetcher "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type GetFileContentFunc func(url, filePath, userName, password string) (string, error)

func GetFileContent(url, filePath, userName, password string) (string, error) {
	r, err := gitFetcher.Clone(memory.NewStorage(), nil, &gitFetcher.CloneOptions{
		Auth: &http.BasicAuth{
			Username: userName,
			Password: password,
		},
		URL: url,
	})
	if err != nil {
		return "", fmt.Errorf("couldn't clone git repository (%s). err: %w", url, err)
	}

	ref, err := r.Head()
	if err != nil {
		return "", fmt.Errorf("couldn't get head commit of git repository (%s). err: %w", url, err)
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return "", fmt.Errorf("couldn't get commit object of ref. err: %w", err)
	}

	file, err := commit.File(filePath)
	if err != nil {
		return "", fmt.Errorf("couldn't get file path (%s). err: %w", filePath, err)
	}

	return file.Contents()
}
