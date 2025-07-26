package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var (
	Banner = ` 


████████╗███████╗██╗  ██╗████████╗██████╗  ██████╗ ██╗  ██╗
╚══██╔══╝██╔════╝╚██╗██╔╝╚══██╔══╝██╔══██╗██╔═══██╗╚██╗██╔╝
   ██║   █████╗   ╚███╔╝    ██║   ██████╔╝██║   ██║ ╚███╔╝ 
   ██║   ██╔══╝   ██╔██╗    ██║   ██╔══██╗██║   ██║ ██╔██╗ 
   ██║   ███████╗██╔╝ ██╗   ██║   ██████╔╝╚██████╔╝██╔╝ ██╗
   ╚═╝   ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═════╝  ╚═════╝ ╚═╝  ╚═╝
                                                           
  `
	purple = lipgloss.Color("99")
	darkBG = lipgloss.Color("#282828")
	style  = lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(purple).
		Foreground(purple).
		PaddingTop(2).
		Width(90).
		Height(31).
		Align(lipgloss.Center)
)

func main() {

	fmt.Println(style.Render(Banner))
}
