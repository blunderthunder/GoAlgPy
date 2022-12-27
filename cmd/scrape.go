package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrape LeetCode challenge Description",
	Long:  "This command will scrape the innerHTML of description section of leetcode and save it as Markdown file.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		selector := "#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div:nth-child(3) > div"

		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()

		// run task and srape
		var res string
		err := chromedp.Run(ctx,
			chromedp.Navigate(args[0]),
			chromedp.Text(selector, &res, chromedp.NodeVisible),
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(strings.TrimSpace(res))
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}
