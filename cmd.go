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
		UserFollower(api)

	case *followingFlag:
		if *userFlag == "" {
			fmt.Println("Error: Username not provided. Use -u flag.")
			os.Exit(1)
		}
		UserFollowing(api)

	case *help:
		Help()
	case *showContribution:
		contributionChart(*userFlag)
	default:
		getUserInfo(api)
	}
}

func UserFollower(api string) {

	getUserInfo(api)
	println(" ____________________________________________________________")
	println("|                                                            |")
	println("|--------------------[ Followers ]---------------------------|")
	getUserFollower(api)

}

func UserFollowing(api string) {
	getUserInfo(api)
	println(" ____________________________________________________________")
	println("|                                                            |")
	println("|--------------------[ Following ]---------------------------|")
	getUserFollowing(api)
}
func Help() {
	println(" ____________________________________________________________________________________________________________________")
	println("|                                                                                                                    |")
	println(style.Render("| The program will then fetch all repositories from that user and print out their names, description, url, etc...|"))
	println("|____________________________________________________________________________________________________________________|")
	println("| 1) To get user details                                                                                             |")
	println("|                                                                                                                    |")
	println("| --> go run main.go -u {git_username}                                                                               |")
	println("|____________________________________________________________________________________________________________________|")
	println("| 2) To get user followers use -F flag                                                                               |")
	println("|                                                                                                                    |")
	println("| --> go run main.go -u {git_username} -F                                                                            |")
	println("|____________________________________________________________________________________________________________________|")
	println("|                                                                                                                    |")
	println("| 3) To get user following use -f flag                                                                               |")
	println("| --> go run main.go -u {git_username} -f                                                                            |")
	println("|____________________________________________________________________________________________________________________|")

}
