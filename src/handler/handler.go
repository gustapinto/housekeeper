package handler

import (
	"log"
	"os"
	"os/exec"
)

func PyHandler(script string) {
	// Executes a python script using a system call on the script path
	out, err := exec.Command("python", script).Output()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(out)
}

func TxtHandler(filePath string) {
	// Read a file content using the os package, this is not so
	// efficient on larger files but is a pratical and high level way
	// of doing it
	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(content))
}
