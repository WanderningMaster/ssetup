package store

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

const APP_FOLDER = "/ssetup"

type Script struct {
	Path string
	Exec string
	Name string
}

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
	found, err := exists(dir)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !found {
		fmt.Printf("Creating Application local dir...\n")
		fmt.Println(dir)
		os.Mkdir(dir, 0750)
	}
}

func MatchExt(fileName string) (string, string, error) {
	execExt := map[string]string{
		"sh": "/bin/bash",
		"js": "/bin/node",
		"py": "/bin/python3",
	}

	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		return "", "", errors.New("invalid filename")
	}
	ext := parts[1]
	exec, found := execExt[ext]
	if !found {
		return "", "", errors.New("unsupported script extension")
	}

	return exec, parts[0], nil
}

func ListScripts() ([]Script, error) {

	dir := GetLocalDataDir()
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files in directory:")
	list := []Script{}
	for _, file := range files {
		exec, name, err := MatchExt(file.Name())
		if err != nil {
			return []Script{}, err
		}
		script := Script{
			Path: dir + "/" + file.Name(),
			Exec: exec,
			Name: name,
		}
		list = append(list, script)
	}

	return list, nil
}
