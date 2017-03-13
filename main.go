package main

import (
	"flag"
	. "github.com/mdelillo/claimer/claimer"
	. "github.com/mdelillo/claimer/fs"
	. "github.com/mdelillo/claimer/git"
	. "github.com/mdelillo/claimer/locker"
	. "github.com/mdelillo/claimer/slack"
	"io/ioutil"
	"os"
)

func main() {
	apiToken := flag.String("apiToken", "", "API Token for Slack")
	repoUrl := flag.String("repoUrl", "", "URL for git repository of locks")
	deployKey := flag.String("deployKey", "", "Deploy key for Github")
	flag.Parse()

	gitDir, err := ioutil.TempDir("", "claimer-git-repo")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(gitDir)

	fs := NewFs()
	repo := NewRepo(*repoUrl, *deployKey, gitDir)
	locker := NewLocker(fs, repo)
	slackClient := NewClient("https://slack.com", *apiToken)
	claimer := New(locker, slackClient)
	claimer.Run()
}
