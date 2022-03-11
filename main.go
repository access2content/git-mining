package main

import (
	"flag"

	"github.com/access2content/go-practice/git-mining/git"
)

func main() {
	//	GET the input path of the git folder as input from flag
	gitPath := flag.String("path", ".", "Path of the folder containing .git folder")
	flag.Parse()

	git.Commits = git.GetAllCommits(*gitPath)
	git.GetCommiterContributions(git.Commits)
}
