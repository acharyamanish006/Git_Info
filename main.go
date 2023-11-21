package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	// "github.com/go-delve/delve/service/api"
)

type GitInfo struct {
	Name      string `json:"name"`
	Url       string `json:"html_url"`
	Bio       string `json:"bio"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Repos     int    `json:"public_repos"`
}

type Follow []struct {
	User string `json:"login"`
}

func main() {
	user := os.Args[1]

	if len(os.Args) < 1 {
		fmt.Println("Username not provided")
		return
	}

	api := "https://api.github.com/users/" + user

	getUserInfo(api)
	println(" ____________________________________________________________")
	println("|                                                            |")
	println("|--------------------[ Followers ]---------------------------|")
	getUserFollower(api)
	println("|--------------------[ Following ]---------------------------|")
	getUserFollowing(api)

}

func getUserInfo(api string) {

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

func getUserFollower(api string) {
	res, err := http.Get(api + "/followers")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var follower Follow

	json.Unmarshal(body, &follower)

	for i := 0; i < len(follower); i++ {
		fmt.Println("|-->", (follower[i].User))

	}

	// fmt.Printf(follower)
}

func getUserFollowing(api string) {
	res, err := http.Get(api + "/following")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var following Follow

	json.Unmarshal(body, &following)
	for i := 0; i < len(following); i++ {
		fmt.Println("|-->", following[i].User)

	}
}
