package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func isRustPresent() error {
	log.Println("Checking the presence of DARK LORD `cargo`  ...")
	_, err := exec.Command("cargo", "help").Output()
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
