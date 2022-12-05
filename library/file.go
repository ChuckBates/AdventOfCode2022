package library

import (
	"log"
	"os"
)

func LoadFile(input string) string {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
