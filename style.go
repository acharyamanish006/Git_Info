package main

import (
	"github.com/charmbracelet/lipgloss"
)

var Center = lipgloss.NewStyle().Align(lipgloss.Center)
var widthFull = lipgloss.NewStyle().Width(100)
var background = lipgloss.NewStyle().Background(lipgloss.Color("33")).Border(lipgloss.RoundedBorder()).Foreground(lipgloss.Color("#FFFFFF"))
var border = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
var whiteColor = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
var Padding = lipgloss.NewStyle().PaddingRight(2).PaddingLeft(2)
var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Width(120)
var alignCenter = lipgloss.NewStyle().Align(lipgloss.Center)
var alignLeft = lipgloss.NewStyle().Align(lipgloss.Left).PaddingLeft(2)
var Bold = lipgloss.NewStyle().Bold(true)
