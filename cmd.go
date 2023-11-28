package main

import (
	"flag"
	"fmt"
	"os"
)

func argsCommand() {
	userFlag := flag.String("u", "", "Specify the username")
	followerFlag := flag.Bool("F", false, "Specify the follower")
	followingFlag := flag.Bool("f", false, "Specify the following")
	help := flag.Bool("h", false, "Specify the repository name")
	showContribution := flag.Bool("c", false, "contribution")

	flag.Parse()

	api := "https://api.github.com/users/" + *userFlag

	switch {
	case *followerFlag:
		if *userFlag == "" {
			fmt.Println("Error: Username not provided. Use -u flag.")
			os.Exit(1)
		}
		UserFollower(api, *userFlag)

	case *followingFlag:
		if *userFlag == "" {
			fmt.Println("Error: Username not provided. Use -u flag.")
			os.Exit(1)
		}
		UserFollowing(api, *userFlag)

	case *help:
		(Help())
	case *showContribution:
		contributionChart(*userFlag)
	default:
		getUserInfo(api)
	}
}

func UserFollower(api string, user string) {

	getUserInfo(api)
	println((FollowStyle.Render("Followers")))
	getUserFollower(api, user)

}

func UserFollowing(api string, user string) {
	getUserInfo(api)
	println((FollowStyle.Render("Following")))
	getUserFollowing(api, user)
}
func Help() {
	helpHeader := " The program will then fetch all repositories from that user and print out their names, description, url, etc...\n"
	help_1 := "\n1) To get user details\n --> go run main.go -u {git_username}\n"

	help_2 := "\n2) To see the followers of a particular git user\n --> go run main.go -u {git_username} -F\n"

	help_3 := "\n3) To get user following use -f flag \n --> go run main.go -u {git_username} -f  \n"

	helpHeader = (helpStyle.Render(alignCenter.Render(Bold.Render(helpHeader))))

	helpOption := help_1 + help_2 + help_3

	helpOption = alignLeft.Render(whiteColor.Render(helpOption))

	help := helpHeader + helpOption

	fmt.Println(border.Render(help))
}
