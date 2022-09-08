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
		selector := "#app > div > div.main__2_tD > div.content__3fR6 > div > div.side-tools-wrapper__1TS9 > div > div.css-1gd46d6-Container.e5i1odf0 > div.css-jtoecv > div > div.tab-pane__ncJk.css-1eusa4c-TabContent.e5i1odf5 > div > div.content__u3I1.question-content__JfgR"

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
