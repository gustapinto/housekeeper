package handler

import (
	"fmt"
	"os"
	"os/exec"
)

func PyHandler(script string) {
	out, _ := exec.Command("python", script).Output()

	fmt.Println(out)
}

func TxtHandler(filePath string) {
	// Read a file content using the os package, this is not so
	// efficient on larger files but is a pratical and high level way
	// of doing it
	content, _ := os.ReadFile(filePath)

	fmt.Println(string(content))
}
