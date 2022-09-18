package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func isValidLeetCodeUrl(url string) error {
	if strings.HasPrefix(url, "https://leetcode.com/problems/") {
		return nil
	}
	return fmt.Errorf("%s Not valid leetcode url", url)
}

func isRustPresent() error {
	log.Println("Checking the presence of DARK LORD `cargo`  ...")
	_, err := exec.Command("cargo", "help").Output()
	return err
}

func isGoCodeAvailable(path string) error {
	_, err := os.ReadFile(fmt.Sprintf("%s/main.go", path))
	return err
}

func isPyCodeAvailable(path string) error {
	_, err := os.ReadFile(fmt.Sprintf("%s/main.py", path))
	return err
}

func isRustCodeAvailable(path string) error {
	_, err := os.ReadFile(fmt.Sprintf("%s/Cargo.toml", path))
	return err
}

func createFile(filename string) (string, error) {
	log.Println("create file  : ", filename)
	_, err := os.Create(filename)
	return filename, err
}

func createDir(foldername string) error {
	log.Println("create Directory : ", foldername)
	err := os.Mkdir(foldername, 0750)
	return err
}

func writeToFile(filename, content string) {
	os.WriteFile(filename, []byte(content), os.ModeAppend)
}

func registerChallange(projectname string) {
	createFile(fmt.Sprintf("%s/.challenge", projectname))
	log.Println("Completed Registering Challenge.")
}

func createActiveChallange(projectname string) {
	os.Remove(".activechallange")
	filename, _ := createFile(".activechallange")
	writeToFile(filename, projectname)
}

func parseActiveProject() string {
	data, err := os.ReadFile(".activechallange")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Fetched active Challenge !")
	return string(data)
}

func executeAndLog(proj string, ptype string, cmd *exec.Cmd) {
	// check for the timespent
	res, timespent, memConsumed, err := runSubprocess(cmd)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("")
	log.Printf("Running %s Code : %s\n", ptype, res)
	log.Printf("Time Spent on %s : %s\n", ptype, timespent)
	log.Printf("Memory Consumed by %s : %d\n", ptype, memConsumed)
	fmt.Println("")

	// save the result to file
	saveExecutedChallengeStat(proj, ptype, timespent, memConsumed)
}

func runSubprocess(cmd *exec.Cmd) (string, time.Duration, int64, error) {

	starttimer := time.Now()
	output, err := cmd.Output()
	elapsed := time.Since(starttimer)

	return string(output), elapsed, cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss, err
}

func saveExecutedChallengeStat(proj string, ptype string, timespent time.Duration, maxMemory int64) {
	statfile := fmt.Sprintf("%s/.%schallenge.stat", proj, ptype)

	content := fmt.Sprintf("timespent:\t%s\t\tmemConsumed:\t%d\n", timespent, maxMemory)

	// check if file already exist
	_, err := os.Open(statfile)

	if os.IsNotExist(err) {
		createFile(statfile)
	}

	writeToFile(statfile, content)
}
