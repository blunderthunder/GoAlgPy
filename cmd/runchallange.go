package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var possibleArgs = []string{"py", "python", "rust", "go", "all"}

func runPyCode(activeProj string) {
	// check python code first
	err := isPyCodeAvailable(activeProj)
	if err != nil {
		log.Println("Python Code not found.")
		return
	}
	executeAndLog(activeProj, "py", exec.Command("python3", fmt.Sprintf("%s/main.py", activeProj)))
}

func runRustCode(activeProj string) {

	err := isRustCodeAvailable(activeProj)
	if err != nil {
		log.Println("Rust Code not found.")
		return
	}
	executeAndLog(activeProj, "rust", exec.Command("cargo", "run", "--manifest-path", fmt.Sprintf("%s/Cargo.toml", activeProj)))
}

func runGoCode(activeProj string) {

	err := isGoCodeAvailable(activeProj)
	if err != nil {
		log.Println("Go Code not found.")
		return
	}
	executeAndLog(activeProj, "go", exec.Command("go", "run", fmt.Sprintf("%s/main.go", activeProj)))
}

func runAllCode(activeProj string) {
	// run python code first
	runPyCode(activeProj)

	// run golang project
	runGoCode(activeProj)

	// run rust project
	runRustCode(activeProj)
}

var runChallengeCmd = &cobra.Command{
	Use:   "run_challenge",
	Short: "run the challenge solution written in different language",
	Long:  "Use this command to check solution of different language ` viable options are [ py , rust, go , all] ` ",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts 1 arg(s), received %d \n\n only one of `%s` args are accepted", len(args), "[ "+strings.Join(possibleArgs, ", ")+" ]")
		}
		for _, val := range possibleArgs {
			if val == args[0] {
				return nil
			}
		}
		return fmt.Errorf("only one of `%s` args are accepted , Instead received `%s` ", "[ "+strings.Join(possibleArgs, ", ")+" ]", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		// get current active challenge
		activeChallenge := parseActiveProject()

		if args[0] == "py" || args[0] == "python" {
			runPyCode(activeChallenge)
		} else if args[0] == "go" {
			runGoCode(activeChallenge)
		} else if args[0] == "rust" {
			runRustCode(activeChallenge)
		} else {
			runAllCode(activeChallenge)
		}

	},
}

func init() {
	rootCmd.AddCommand(runChallengeCmd)
}
