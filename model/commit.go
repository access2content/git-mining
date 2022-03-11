package model

type Commit struct {
	CommitHash string
	Committer  string
	Date       string
	Subject    string
	Body       string
	Files      []string
}

type Contribution struct {
	Committer string
	Commits   int
}

type Contributions []Contribution

//	Operations on list of contributions to use SORT function
func (c Contributions) Len() int           { return len(c) }
func (c Contributions) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Contributions) Less(i, j int) bool { return c[i].Commits < c[j].Commits }
