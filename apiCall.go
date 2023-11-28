package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

	user := green.Render(name) + "\n" + yellow.Render(location) + "\n" + eal.Render(bio) + "\n" + red.Render(repos) + "\n" + orange.Render(follower) + "\n" + blue.Render(following)

	fmt.Println(border.Render(Padding.Render(user)))

	// fmt.Printf(("Name: %s \nUrl: %s \nBio: %s \nLocation: %s \nRepos: %d \nFollowers: %d \nFollowing: %d  \n", info.Name, info.Url, info.Bio, info.Location, info.Repos, info.Followers, info.Following))
}

func getUserFollower(api string, username string) {
	res, err := http.Get(api + "/followers")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var follower Follow

	json.Unmarshal(body, &follower)

	var user []string
	// fmt.Printf("%q has %d Followers\n", username, len(follower))

	for i := 0; i < len(follower); i++ {
		// fmt.Println("|-->", follower[i].User)
		user = append(user, follower[i].User)
	}
	Select(user, "Followers", username)
}

func getUserFollowing(api string, username string) {
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
	// fmt.Printf("%q has %d Following\n", username, len(following))

	for i := 0; i < len(following); i++ {
		// fmt.Println("|-->", follower[i].User)
		user = append(user, following[i].User)
	}
	// data := user

	Select(user, "Following", username)
}

func Select(user []string, follow string, username string) {

	row := int(math.Ceil(float64(len(user)) / 4.0))

	data := create2DArray(user, (row), 4)

	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("#929292"))).
		Width(72).
		Rows(data...).
		StyleFunc(func(row, col int) lipgloss.Style {

			even := row%2 == 0

			if even {
				return baseStyle.Copy().Foreground(lipgloss.Color("#00E2C7")).Bold(true).PaddingRight(1).PaddingLeft(1)
			}
			return baseStyle.Copy().Foreground(lipgloss.Color("#FDFF90")).Bold(true).PaddingRight(1).PaddingLeft(1)
		})
	fmt.Println(t)

	// prompt := promptui.Select{
	// 	Label: ("List of " + (username) + " " + follow),
	// 	Items: t,
	// 	Size:  len(user),
	// }

	// _, result, err := prompt.Run()
	// // fmt.Printf("%q: has %d %s\n", result, len(user), follow)
	// FollowStyle.Render()
	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// fmt.Printf("Showing The Information of:%q\n", result)
	// // fmt.Printf("%q: has %d %s\n", result, len(user), follow)

	// api := "https://api.github.com/users/" + result
	// getUserInfo(api)
}

// func create2DArray(stringsArray []string) [][]string {
// 	var array2D [][]string

// 	// Create an array of numbers with the same length as the array of strings
// 	numbersArray := make([]int, len(stringsArray))
// 	for i := range numbersArray {
// 		numbersArray[i] = i + 1
// 	}

// 	// Populate the 2D array
// 	for i := 0; i < len(numbersArray); i++ {
// 		row := []string{fmt.Sprint(numbersArray[i]), stringsArray[i]}
// 		array2D = append(array2D, row)
// 	}

// 	return array2D
// }
func create2DArray(data []string, rows, cols int) [][]string {
	var array2D [][]string

	// // Check if there are enough elements for the specified rows and cols
	// if rows*cols > len(data) {
	// 	fmt.Println("Not enough elements to create a 2D array.")
	// 	return nil
	// }

	// Populate the 2D array
	for i := 0; i < rows; i++ {
		start := i * cols
		end := (i + 1) * cols

		array2D = append(array2D, data[start:end])
	}

	return array2D
}
