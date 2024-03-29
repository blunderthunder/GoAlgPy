package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func scrapeDesc(url string) string {
	selector := "#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div:nth-child(3) > div"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task and srape
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.InnerHTML(selector, &res, chromedp.NodeVisible),
	)

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(res)
}

func newProjDir(projectName, url string) error {

	// check if cargo exist
	if isRustPresent() != nil {
		log.Println("Cargo not present! ")
		// create dir yourself here
		os.Mkdir(projectName, 0750)
	} else {
		_, err := exec.Command("cargo", "new", projectName).Output()
		if err != nil {
			return err
		}
		// create test directory first
		createDir(fmt.Sprintf("%s/tests", projectName))
		// create rust test files
		createFile(fmt.Sprintf("%s/tests/%s_test.rs", projectName, projectName))
		createFile(fmt.Sprintf("%s/src/lib.rs", projectName))
	}
	// create a necessary files
	filename, _ := createFile(fmt.Sprintf("%s/readme.md", projectName))
	writeToFile(filename, scrapeDesc(url))

	createFile(fmt.Sprintf("%s/__init__.py", projectName))
	createFile(fmt.Sprintf("%s/main.py", projectName))
	createFile(fmt.Sprintf("%s/main.go", projectName))

	filename, _ = createFile(fmt.Sprintf("%s/solution.md", projectName))
	writeToFile(filename, fmt.Sprintf("# Solution of %s \n", strings.ReplaceAll(projectName, "_", " ")))
	// set activechallange state to newly created project
	createActiveChallange(projectName)

	log.Printf("Successfully created Project %s \n", strings.ReplaceAll(projectName, "_", " "))
	registerChallange(projectName)
	return nil
}

var neoChallangeCmd = &cobra.Command{
	Use:   "create_challenge",
	Short: "create new directory with the challange",
	Long:  "This command will create new project directory for leetcode challange with all required setting and file in place",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		prompt := promptui.Prompt{
			Label:    "Enter LeetCode Challenge URL",
			Validate: isValidLeetCodeUrl,
		}

		parsedUrl, err := prompt.Run()

		if err != nil {
			log.Fatalf("Prompt failed %s\n", err)
		}
		urlSplit := strings.Split(strings.TrimSpace(parsedUrl), "/")

		projectName := urlSplit[len(urlSplit)-1]
		if strings.TrimSpace(projectName) == "" {
			projectName = urlSplit[len(urlSplit)-2]
		}

		projectName = strings.TrimSpace(strings.ReplaceAll(projectName, "-", "_"))

		// create project related files and directory here
		newProjDir(projectName, parsedUrl)
	},
}

func init() {
	rootCmd.AddCommand(neoChallangeCmd)
}
