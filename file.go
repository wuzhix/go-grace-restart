package restart

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	pidPath = "./go-grace-restart.pid"
)

func getPid() int {
	file, err := os.Open(pidPath)
	defer file.Close()
	if err != nil {
		return 0
	} else {
		buffer := make([]byte, 10)
		count, err := file.Read(buffer)
		if err != nil {
			return 0
		}
		pid, err := strconv.Atoi(string(buffer[:count]))
		if err != nil {
			return 0
		} else {
			if exist(pid) {
				return pid
			} else {
				return 0
			}
		}
	}
}

func savePid() {
	file, err := os.OpenFile(pidPath, os.O_WRONLY|os.O_CREATE, 0777)
	defer file.Close()
	if err == nil {
		_, err = file.Write([]byte(strconv.Itoa(os.Getpid())))
	}
}

func findFile(pid int) string {
	var path string
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/C", "wmic process where processid=" + strconv.Itoa(pid) + " get executablepath")
		ret, err := cmd.Output()
		if err == nil {
			retStr := strings.Replace(string(ret), " ", "", -1)
			retStr = strings.Replace(retStr, "\r", "", -1)
			retArr := strings.Split(retStr, "\n")
			path = retArr[1]
		}
	case "linux":
		//cmd := exec.Command("bash", "-c", `ls -l /proc/`+strconv.Itoa(pid)+` | grep exe | awk -F " " '{print $NF}'`)
		cmd := exec.Command("bash", "-c", `ls -l /proc/`+strconv.Itoa(pid)+` | grep exe`)
		ret, err := cmd.Output()
		if err == nil {
			retArr := strings.Split(strings.Replace(string(ret), "\n", "", -1), " ")
			for i := 0; i < len(retArr); i++ {
				if retArr[i] == "->" {
					path = retArr[i+1]
				}
			}
		}
	default:
	}
	return path
}

func deleteFile(file string) error {
	err := os.Remove(file)
	return err
}
