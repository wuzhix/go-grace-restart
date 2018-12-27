package restart

import (
	"os/exec"
	"strconv"
	"strings"
)

func findFile(pid int) string {
	var path string
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
	return path
}
