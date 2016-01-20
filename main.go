package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("You need to provide some command")
		os.Exit(1)
	}
	command := os.Args[1]
	args := os.Args[2:]
	out, err := runCmd(command, args...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(out))
}

func runCmd(command string, args ...string) ([]byte, error) {
	defer timeTrack(time.Now())
	out, err := exec.Command(command, args...).Output()
	return out, err
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time=%fs \n", elapsed.Seconds())
}
