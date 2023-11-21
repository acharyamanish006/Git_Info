package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	// "github.com/go-delve/delve/service/api"
)

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
	// user := os.Args[1]

	// if len(os.Args) < 1 {
	// 	fmt.Println("Username not provided")
	// 	return
	// }
	argsCommand()

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

	fmt.Printf("Name: %s \nUrl: %s \nBio: %s \nLocation: %s \nRepos: %d \nFollowers: %d \nFollowing: %d  \n", info.Name, info.Url, info.Bio, info.Location, info.Repos, info.Followers, info.Following)
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
	println("| The program will then fetch all repositories from that user and print out their names, description, url, etc...    |")
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

type GitHubEvent struct {
	Type       string `json:"type"`
	CreatedAt  string `json:"created_at"`
	Repository struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func contributionChart(username string) {
	events, err := fetchGitHubEvents(username)
	if err != nil {
		fmt.Println("Error fetching GitHub events:", err)
		return
	}

	contributions := getContributions(events)
	displayContributionChart(contributions, 7)
}

func fetchGitHubEvents(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var events []GitHubEvent
	err = json.NewDecoder(response.Body).Decode(&events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func getContributions(events []GitHubEvent) []bool {
	// Get the current day of the week
	currentDay := int(time.Now().Weekday())

	// Initialize contributions slice
	contributions := make([]bool, 7)

	// Update contributions based on GitHub events
	for _, event := range events {
		// Check if the event is a push event
		if event.Type == "PushEvent" {
			// Get the day of the week for the event
			createdAt, err := time.Parse(time.RFC3339, event.CreatedAt)
			if err == nil {
				dayOfWeek := int(createdAt.Weekday())
				// Update the contribution status for that day
				contributions[(dayOfWeek-currentDay+7)%7] = true
			}
		}
	}

	return contributions
}

func displayContributionChart(contributions []bool, width int) {
	// Display the contribution chart
	fmt.Printf(" Contribution Chart \n\n")
	fmt.Println(" __________________________________")
	fmt.Println("|    |    |    |    |    |    |    |")
	fmt.Println("| Mo | Tu | We | Th | Fr | Sa | Su |")
	fmt.Println("|____|____|____|____|____|____|____|")

	fmt.Println(" __________________________________")
	fmt.Println("|    |    |    |    |    |    |    |")

	for i := 0; i < len(contributions); i++ {
		// Display the day of the week
		// if i%7 == 0 {
		// 	// fmt.Printf("\n")
		// }

		// Display the contribution status
		if contributions[i] {
			fmt.Print("| X  ")
		} else {
			fmt.Print("| _  ")
		}
	}
	fmt.Print("|")
	fmt.Println("\n|____|____|____|____|____|____|____|")

	fmt.Println("\n\n X: Contribution,   Blank: No Contribution")
}
