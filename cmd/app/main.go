package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Fu1cr4m/PortSage/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// 检查是否为 Windows
	if runtime.GOOS != "windows" {
		fmt.Println("PortSage is currently designed for Windows Only.")
		os.Exit(1)
	}

	// 启动 UI
	p := tea.NewProgram(ui.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}