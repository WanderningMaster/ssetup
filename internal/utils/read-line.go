package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input := strings.TrimRight(scanner.Text(), "\r\n")
	return input
}
