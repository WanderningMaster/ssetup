package store

import (
	"fmt"
	"log"
	"os"
)

const APP_FOLDER = "/ssetup"

func GetTempDir() string {
	return os.TempDir() + APP_FOLDER
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
	dir := GetTempDir()
	found, err := exists(dir)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !found {
		fmt.Printf("Creating Application temp dir...")
		os.Mkdir(dir, 0777)
	}
}
