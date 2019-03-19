package main

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	excuteShell("mkdir aabbccxxx")

}

func excuteShell(commond string) bool {
	c := commond
	var cmd *exec.Cmd
	if runtime.GOOS != "windows" {
		cmd = exec.Command("bash", "-c", c)
	} else {
		cmd = exec.Command("cmd.exe", "/c", c)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
		return false
	}
	return true
}
