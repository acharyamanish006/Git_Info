package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/manifoldco/promptui"
)

func getUserInfo(api string) {

	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var info GitInfo

	// var user string;

	json.Unmarshal(body, &info)

	name := "Name: " + info.Name
	// email:= "Email: " + info.Email
	// company:= "Company: " + info.Company
	location := "Location: " + info.Location
	bio := "Bio: " + info.Bio
	repos := "Repos: " + strconv.Itoa(info.Repos)
	follower := "Followers: " + strconv.Itoa(info.Followers)
	following := "Following: " + strconv.Itoa(info.Following)

	user := name + "\n" + location + "\n" + bio + "\n" + repos + "\n" + follower + "\n" + following

	fmt.Println(border.Render(Padding.Render(user)))

	// fmt.Printf(("Name: %s \nUrl: %s \nBio: %s \nLocation: %s \nRepos: %d \nFollowers: %d \nFollowing: %d  \n", info.Name, info.Url, info.Bio, info.Location, info.Repos, info.Followers, info.Following))
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

	var user []string

	for i := 0; i < len(follower); i++ {
		// fmt.Println("|-->", follower[i].User)
		user = append(user, follower[i].User)
	}
	Select(user)
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
	// for i := 0; i < len(following); i++ {
	// 	fmt.Println("|-->", following[i].User)
	// }
	var user []string

	for i := 0; i < len(following); i++ {
		// fmt.Println("|-->", follower[i].User)
		user = append(user, following[i].User)
	}
	Select(user)
}

func Select(user []string) {
	prompt := promptui.Select{
		Label: "Select User",
		Items: user,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
	api := "https://api.github.com/users/" + result
	getUserInfo(api)
}
