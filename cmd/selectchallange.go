package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func listOutProjects() []string {
	current_dir, _ := os.Getwd()
	f, _ := os.Open(current_dir)
	files, _ := f.Readdir(0)

	var listDir []string

	for _, file := range files {
		if file.IsDir() {
			_, err := os.ReadFile(fmt.Sprintf("%s/%s/.challenge", current_dir, file.Name()))
			if err != nil {
				continue
			} else {
				listDir = append(listDir, file.Name())
			}
		}
	}
	return listDir
}

var selectChallengeCmd = &cobra.Command{
	Use:   "select_challenge",
	Short: "select challenge to flag it as active project",
	Long:  "Used to select challenge as a active challenge from list of project",

	Args: cobra.ExactArgs(0),

	Run: func(cmd *cobra.Command, args []string) {

		projects := listOutProjects()

		selectPromt := promptui.Select{
			Label: "Choose Any of the Below challenge to set as Active Challenge.",
			Items: projects,
			Size:  20,
		}

		_, result, err := selectPromt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		createActiveChallange(result)

		log.Printf("Successfully Switched to %s challenge.\n", result)
	},
}

func init() {
	rootCmd.AddCommand(selectChallengeCmd)
}
