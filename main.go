package main

type GitInfo struct {
	Name      string `json:"name"`
	Url       string `json:"html_url"`
	Bio       string `json:"bio"`
	Location  string `json:"location"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Repos     int    `json:"public_repos"`
}

type Follow []struct {
	User string `json:"login"`
}

func main() {
	ArgsCommand()
}
