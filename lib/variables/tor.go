package variables

import (
	"fmt"
	"os/exec"
	"time"
)




var (
	LastIP string
)


func ReloadTor() {
	for {
		Reload()
		time.Sleep(3*time.Second)
	}
}

func Reload() {
	_, err := exec.Command("systemctl", "reload", "tor").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	out, err := exec.Command("torsocks", "curl", "ipv4.icanhazip.com").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	if string(out) == LastIP {
		Reload()
	}
	LastIP = string(out)
}