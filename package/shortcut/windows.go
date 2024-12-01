package shortcut

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

func OpenBrowser(url string) {
	cmd := exec.Command("cmd", "/c", "start", url)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
func RunCommandLine(command string) {
	args := strings.Split(command, " ")
	args = append([]string{"/c"}, args...)
	cmd := exec.Command("cmd", args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
