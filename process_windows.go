package restart

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func kill(pid int) error {
	process, err := os.FindProcess(pid)
	if err == nil {
		return process.Kill()
	}
	return err
}

//判断进程是否存在
func exist(pid int) bool {
	ret := false
	//cmd := exec.Command("cmd", "/C", `tasklist /fi \"pid eq 360\" | findstr 360`)
	cmd := exec.Command("cmd", "/C", "tasklist", "/fi", "pid eq "+strconv.Itoa(pid))
	out, err := cmd.Output()
	if err == nil {
		outArr := strings.Split(string(out), "\r\n")
		for i := 0; i < len(outArr); i++ {
			if strings.Contains(outArr[i], " "+strconv.Itoa(pid)+" ") {
				ret = true
				break
			}
		}
	}
	return ret
}
