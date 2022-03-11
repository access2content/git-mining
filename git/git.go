package git

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"sort"
	"strings"

	"github.com/access2content/go-practice/git-mining/model"
)

var Commits []model.Commit

//	GET all git commits from a .git folder
func GetAllCommits(path string) []model.Commit {

	//	Run the git log command on that folder
	cmd := exec.Command("git", "--git-dir", path+"/.git", "log", `--pretty=format:^^^^^{"CommitHash": "%H", "Committer":"%aN", "Date": "%at"}^^^%n%s^^^%n%b^^^`, "--name-only")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running git command")
	}

	//	SPLIT the commits
	allCommits := strings.Split(string(output), "^^^^^")

	newCommit := &model.Commit{}
	commitList := []model.Commit{}
	//	LOOP through all the commits and extract only the JSON object first
	for i := len(allCommits) - 1; i > 0; i-- {
		commitDetails := strings.Split(allCommits[i], "^^^\n")

		//	GET JSON details and store in struct
		commitJSON := commitDetails[0]
		json.Unmarshal([]byte(commitJSON), newCommit)

		//	Save the Commit Subject
		commitSubject := commitDetails[1]
		newCommit.Subject = commitSubject

		//	Save the Commit Body
		commitBody := commitDetails[2]
		newCommit.Body = commitBody

		//	Save the Commit Files
		commitFiles := strings.Split(commitDetails[3], "\n")
		fileList := []string{}
		for j := 0; j < len(commitFiles)-2; j++ {
			fileList = append(fileList, commitFiles[j])
		}
		newCommit.Files = fileList

		commitList = append(commitList, *newCommit)
	}

	return commitList
}

//	Get all Committers
func GetCommiterContributions(allCommits []model.Commit) map[string]int {
	//	POPULATE the contribution
	committers := make(map[string]int)
	for i := 0; i < len(allCommits); i++ {
		_, exists := committers[allCommits[i].Committer]
		if exists {
			committers[allCommits[i].Committer]++
		} else {
			committers[allCommits[i].Committer] = 1
		}
	}

	//	CONVERT it to an array of struct
	contributions := model.Contributions{}
	for committer, commits := range committers {
		contributions = append(contributions, model.Contribution{
			Committer: committer,
			Commits:   commits,
		})
	}

	//	SORT the array
	sort.Sort(contributions)

	//	PRINT contributions in ascending Order
	for _, k := range contributions {
		fmt.Printf("%v: %v\n", k.Committer, k.Commits)
	}
	return committers
}
