/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// scrollCmd represents the scroll command
var scrollCmd = &cobra.Command{
	Use:   "scroll",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scroll called")
		startScrollingText()
	},
}

func init() {
	rootCmd.AddCommand(scrollCmd)
}

func startScrollingText() {
	textBlock := []string{
		"Press Any Key To Unlock",
	}
	scrollSpeed := time.Millisecond * 150

	for {
		// Create a temporary slice to hold the updated lines
		updatedTextBlock := make([]string, len(textBlock))

		// Update each line in the temporary slice
		for i, line := range textBlock {
			if len(line) > 1 {
				updatedTextBlock[i] = line[1:] + line[:1]
			} else {
				updatedTextBlock[i] = line
			}
		}

		// Copy the updated lines back to the original textBlock
		copy(textBlock, updatedTextBlock)

		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Print the entire updated block at once
		for _, line := range updatedTextBlock {
			fmt.Println(line)
		}

		// Pause for the scroll speed duration
		time.Sleep(scrollSpeed)
	}
}
