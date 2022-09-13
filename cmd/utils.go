package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
		log.Fatal()
	}
	log.Println("Fetched active Challenge !")
	return string(data)
}
