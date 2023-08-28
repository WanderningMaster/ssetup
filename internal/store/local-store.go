package store

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

const APP_FOLDER = "/ssetup"

func GetLocalDataDir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username
	return fmt.Sprintf("/home/%s/Documents", username) + APP_FOLDER
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SetupStore() {
	dir := GetLocalDataDir()
	// fmt.Printf("Local dir: %s\n", dir)
	found, err := exists(dir)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !found {
		fmt.Printf("Creating Application local dir...\n")
		fmt.Println(dir)
		os.Mkdir(dir, 0750)
		// os.Mkdir(dir, 0777)
	}
}
