package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func contributionChart(username string) {
	re := lipgloss.NewRenderer(os.Stdout)
	labelStyle := re.NewStyle().Foreground(lipgloss.Color("241"))

	board := generateContributionBoard(7, 53)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		Rows(board...).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 1)
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
