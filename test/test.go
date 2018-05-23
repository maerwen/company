package main

import "os/exec"

func main() {
	cmd := exec.Command("/opt/google/chrome")
	cmd.Run()
}
