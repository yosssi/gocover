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

	// Execute `go test`.
	execCmd(exec.Command("go", "test", fmt.Sprintf("-coverprofile=%s", file)))

	// Execute `go tool cover`.
	execCmd(exec.Command("go", "tool", "cover", fmt.Sprintf("-html=%s", file)))
}

func execCmd(cmd *exec.Cmd) {
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}
