package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Set the coverage profile file path.
	file := os.TempDir() + "coverage.out"

	var path string

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// Execute `go test`.
	execCmd(exec.Command("go", "test", fmt.Sprintf("-coverprofile=%s", file), path))

	// Execute `go tool cover`.
	execCmd(exec.Command("go", "tool", "cover", fmt.Sprintf("-html=%s", file)))
}

// execCmd executes the specified command.
func execCmd(cmd *exec.Cmd) {
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}
