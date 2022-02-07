package Function

import "os/exec"

func StartNmap(target string) {
	command := exec.Command("nmap", "-p-", target, "Pn")

	command.Run()
}

func WebScan(target string) {

}
