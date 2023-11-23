// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// type GitHubEvent struct {
// 	Type       string    `json:"type"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	Repository struct {
// 		Name string `json:"name"`
// 	} `json:"repo"`
// }

// func contributionChart(username string) {
// 	events, err := fetchGitHubEvents(username)
// 	if err != nil {
// 		fmt.Println("Error fetching GitHub events:", err)
// 		return
// 	}

// 	contributions := getContributions(events)
// 	displayContributionChart(contributions, 7)
// }

// func fetchGitHubEvents(username string) ([]GitHubEvent, error) {
// 	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
// 	response, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer response.Body.Close()

// 	var events []GitHubEvent
// 	err = json.NewDecoder(response.Body).Decode(&events)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return events, nil
// }

// func getContributions(events []GitHubEvent) []bool {
// 	// Get the current day of the week
// 	currentDay := int(time.Now().Weekday())

// 	// Initialize contributions slice
// 	contributions := make([]bool, 7)

// 	// Update contributions based on GitHub events
// 	for _, event := range events {
// 		// Check if the event is a contribution event
// 		if isContributionEvent(event.Type) {
// 			// Get the day of the week for the event
// 			dayOfWeek := int(event.CreatedAt.Weekday())
// 			// Update the contribution status for that day
// 			contributions[(dayOfWeek-currentDay+7)%7] = true
// 		}
// 	}

// 	return contributions
// }

// func isContributionEvent(eventType string) bool {
// 	// List of contribution event types
// 	contributionEventTypes := map[string]bool{
// 		"PushEvent":        true,
// 		"PullRequestEvent": true,
// 		"IssuesEvent":      true,
// 		// Add more event types as needed
// 	}

// 	return contributionEventTypes[eventType]
// }

// // dayOfWeek := int(createdAt.Weekday())
// //
// func displayContributionChart(contributions []bool, width int) {
// 	// Display the contribution chart
// 	fmt.Printf(" Contribution Chart \n\n")
// 	fmt.Println(" __________________________________")
// 	fmt.Println("|    |    |    |    |    |    |    |")
// 	fmt.Println("| Mo | Tu | We | Th | Fr | Sa | Su |")
// 	fmt.Println("|____|____|____|____|____|____|____|")

// 	fmt.Println(" __________________________________")
// 	fmt.Println("|    |    |    |    |    |    |    |")

// 	for i := 0; i < len(contributions); i++ {
// 		// Display the day of the week
// 		// if i%7 == 0 {
// 		// 	// fmt.Printf("\n")
// 		// }

// 		// Display the contribution status
// 		if contributions[i] {
// 			fmt.Print("| X  ")
// 		} else {
// 			fmt.Print("| _  ")
// 		}
// 	}
// 	fmt.Print("|")
// 	fmt.Println("\n|____|____|____|____|____|____|____|")

// 	fmt.Println("\n\n X: Contribution,   Blank: No Contribution")
// }
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GitHubEvent struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

func contributionChart(username string) {
	events, err := fetchGitHubEventsForYear(username)
	if err != nil {
		fmt.Println("Error fetching GitHub events:", err)
		return
	}

	contributions := getContributions(events)
	displayMonthlyContributionChart(contributions)

}

func fetchGitHubEventsForYear(username string) ([]GitHubEvent, error) {
	var allEvents []GitHubEvent

	// GitHub Events API provides events for the last 90 days by default
	// Make paginated requests to cover the entire year
	for i := 1; i <= 4; i++ {
		events, err := fetchGitHubEventsPage(username, i)
		if err != nil {
			return nil, err
		}
		allEvents = append(allEvents, events...)
	}

	return allEvents, nil
}

func fetchGitHubEventsPage(username string, page int) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events?page=%d", username, page)
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

func getContributions(events []GitHubEvent) map[int]map[time.Month][]int {
	// Initialize contributions map
	contributions := make(map[int]map[time.Month][]int)

	// Update contributions based on GitHub events
	for _, event := range events {
		// Check if the event is a contribution event
		if isContributionEvent(event.Type) {
			// Extract the year, month, and day of the event
			year, month, day := event.CreatedAt.Date()
			// Create a new year entry if it doesn't exist
			if _, ok := contributions[year]; !ok {
				contributions[year] = make(map[time.Month][]int)
			}
			// Create a new month entry if it doesn't exist
			if _, ok := contributions[year][month]; !ok {
				contributions[year][month] = make([]int, 31)
			}
			// Update the day entry
			contributions[year][month][day-1] = 1
		}
	}

	return contributions
}

func isContributionEvent(eventType string) bool {
	// List of contribution event types
	contributionEventTypes := map[string]bool{
		"PushEvent":        true,
		"PullRequestEvent": true,
		"IssuesEvent":      true,
		// Add more event types as needed
	}

	return contributionEventTypes[eventType]
}

func displayMonthlyContributionChart(contributions map[int]map[time.Month][]int) {
	// Display contributions for each month
	for year, monthlyContributions := range contributions {
		for month, dailyContributions := range monthlyContributions {
			fmt.Printf("Contributions for %s %d:\n", month, year)
			displayDailyContributionChart(dailyContributions)
			fmt.Println("----------------------------")
		}
	}
}

func displayDailyContributionChart(dailyContributions []int) {
	// Display the daily contribution chart
	fmt.Println("  Mo Tu We Th Fr Sa Su")
	for i := 0; i < len(dailyContributions); i++ {
		// Display the day of the month
		if i%7 == 0 {
			fmt.Printf("\n-")
		}
		// Display the contribution status
		if dailyContributions[i] == 1 {
			fmt.Print("  X")
		} else {
			fmt.Print("  _")
		}
	}
	fmt.Println()
}
