package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitInfo struct {
	Name      string `json:"name"`
	Url       string `json:"html_url"`
	Bio       string `json:"bio"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Repos     int    `json:"public_repos"`
}

func main() {
	getUserInfo()
}

func getUserInfo() {
	api := "https://api.github.com/users/acharyamanish006"

	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var info GitInfo

	json.Unmarshal(body, &info)

	fmt.Printf("Name: %s \nUrl: %s \nBio: %s \nRepos: %d \nFollowers: %d \nFollowing: %d  \n", info.Name, info.Url, info.Bio, info.Repos, info.Followers, info.Following)
}
