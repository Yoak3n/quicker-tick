package shortcut

import "os/exec"

func OpenBrowser(url string) {
	go func() {
		command := exec.Command("cmd", "/c", "start", url)
		command.Run()
	}()
}
