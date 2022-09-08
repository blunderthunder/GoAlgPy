package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var possibleArgs = []string{"py", "rust", "go", "all"}

var runChallengeCmd = &cobra.Command{
	Use:   "run_challenge",
	Short: "run the challenge solution written in different language",
	Long:  "Use this command to check solution of different language ` viable options are [ py , rust, go , all] ` ",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts 1 arg(s), received %d", len(args))
		}
		for _, val := range possibleArgs {
			if val == args[0] {
				return nil
			}
		}
		return fmt.Errorf("only one of `%s` args are accepted , Instead received `%s` ", "[ "+strings.Join(possibleArgs, ", ")+" ]", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(runChallengeCmd)
}
