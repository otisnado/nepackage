package services

import (
	"context"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/otisnado/nepackage/utils"
	log "github.com/sirupsen/logrus"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GitHubClient() *github.Client {
	var ctx = context.Background()
	var ts = oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	var tc = oauth2.NewClient(ctx, ts)
	var client = github.NewClient(tc)
	return client
}

func CreateGitHubRemoteRepository(repositoryName string, isPrivate bool) (repositoryUrl string, err error) { // Create a repository model
	ctx := context.Background()
	repository := &github.Repository{
		Name:    &repositoryName,
		Private: github.Bool(isPrivate),
	}
	createdRepo, _, err := GitHubClient().Repositories.Create(ctx, "", repository)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	log.Debug("repository created at: ", *createdRepo.HTMLURL)

	return *createdRepo.HTMLURL, nil
}

func PushLocalRepositoryToGitHub(path string, remoteUrl string) (err error) { // Create a repository model
	systemUsername, err := utils.GetCurrentUser()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	repository, err := git.PlainOpen(path)
	if err != nil {
		log.Error("error opening repository ", err.Error())
		return err
	}

	_, err = repository.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteUrl},
	})
	if err != nil {
		log.Error("error adding remote to repository ", err.Error())
		return err
	}

	err = repository.Push(&git.PushOptions{
		RemoteName: "origin",
		RemoteURL:  remoteUrl,
		Auth: &http.BasicAuth{
			Username: systemUsername,
			Password: os.Getenv("GH_TOKEN"),
		},
	})
	if err != nil {
		log.Error("error pushing repository ", err.Error())
		return err
	}

	return nil
}
