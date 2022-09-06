package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrape LeetCode challenge Description",
	Long:  "This command will scrape the innerHTML of description section of leetcode and save it as Markdown file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scraping the file \nScraped Thank you for your patience.")
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}
