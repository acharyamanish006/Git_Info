package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func contributionChart(username string) {
	re := lipgloss.NewRenderer(os.Stdout)
	labelStyle := re.NewStyle().Foreground(lipgloss.Color("245"))

	board := generateContributionBoard(7, 53)

	t := table.New().
		// Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("#91B39F"))).
		BorderRow(true).
		BorderColumn(true).
		Rows(board...).
		StyleFunc(func(row, col int) lipgloss.Style {
			even := rand.Intn(5)
			if even <= 2 {
				return lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#078203"))
			} else {
				// Styling for rows divisible by 3
				return lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#034f01"))
			}

		})

	months := labelStyle.Render(strings.Join([]string{"     January", "     February", "     March", "     April", "     May", "     June", "     July", "     August", "     September", "     October", "     November", "     December     "}, "  "))
	days := labelStyle.Render(strings.Join([]string{" Sun", "Mon", "Tus", "Wes", "Thu", "Fri", "Sat"}, "\n\n "))

	fmt.Println(lipgloss.JoinVertical(lipgloss.Right, lipgloss.JoinHorizontal(lipgloss.Center, days, t.Render()), months) + "\n")

}

func generateContributionBoard(rows, cols int) [][]string {
	board := make([][]string, rows)
	for i := range board {
		board[i] = make([]string, cols)
	}

	return board
}
