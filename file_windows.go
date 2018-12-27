package restart

import (
	"os/exec"
	"strconv"
	"strings"
)

func findFile(pid int) string {
	var path string
	cmd := exec.Command("cmd", "/C", "wmic process where processid=" + strconv.Itoa(pid) + " get executablepath")
	ret, err := cmd.Output()
	if err == nil {
		retStr := strings.Replace(string(ret), " ", "", -1)
		retStr = strings.Replace(retStr, "\r", "", -1)
		retArr := strings.Split(retStr, "\n")
		path = retArr[1]
	}
	return path
}
