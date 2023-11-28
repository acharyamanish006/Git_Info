package main

import (
	"github.com/charmbracelet/lipgloss"
)

var Center = lipgloss.NewStyle().Align(lipgloss.Center)
var widthFull = lipgloss.NewStyle().Width(100)

// var background = lipgloss.NewStyle().Background(lipgloss.Color("33")).Border(lipgloss.RoundedBorder()).Foreground(lipgloss.Color("#FFFFFF"))
var border = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#01BE85")).Width(70)
var whiteColor = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
var Padding = lipgloss.NewStyle().PaddingRight(2).PaddingLeft(2)
var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Width(120)
var alignCenter = lipgloss.NewStyle().Align(lipgloss.Center)
var alignLeft = lipgloss.NewStyle().Align(lipgloss.Left).PaddingLeft(2)
var Bold = lipgloss.NewStyle().Bold(true)

var FollowStyle = lipgloss.NewStyle().Width(70).Border(lipgloss.RoundedBorder()).Foreground(lipgloss.Color("#D7FF87")).Align(lipgloss.Center).Bold(true).BorderForeground(lipgloss.Color("#01BE85"))

// table := table.New().Headers().Border(lipgloss.NormalBorder()).BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
// Width(80).
// Rows(data...).

var green = lipgloss.NewStyle().Foreground((lipgloss.Color("#01BE85")))
var yellow = lipgloss.NewStyle().Foreground((lipgloss.Color("#FDFF90")))
var red = lipgloss.NewStyle().Foreground((lipgloss.Color("#FF7698")))
var orange = lipgloss.NewStyle().Foreground((lipgloss.Color("#FF875F")))
var eal = lipgloss.NewStyle().Foreground((lipgloss.Color("#00E2C7")))
var blue = lipgloss.NewStyle().Foreground((lipgloss.Color("#7D5AFC")))

// "Bug":      lipgloss.Color("#D7FF87"),
// "Electric": lipgloss.Color("#FDFF90"),
// "Fire":     lipgloss.Color("#FF7698"),
// "Flying":   lipgloss.Color("#FF87D7"),
// "Grass":    lipgloss.Color("#75FBAB"),
// "Ground":   lipgloss.Color("#FF875F"),
// "Normal":   lipgloss.Color("#929292"),
// "Poison":   lipgloss.Color("#7D5AFC"),
// "Water":    lipgloss.Color("#00E2C7")
