package restart

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
)

// 杀死进程
func kill(pid int) error {
	files, err := ioutil.ReadDir(procPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			curPid, err := strconv.Atoi(file.Name())
			if err != nil {
				//return err
				println(err.Error())
			} else {
				stat, err := os.Open(procPath + "/" + file.Name() + "/stat")
				if err != nil {
					return err
				} else {
					buffer := make([]byte, 1024)
					stat.Read(buffer)
					stat.Close()
					statArr := strings.Split(string(buffer), " ")
					parentPid, err := strconv.Atoi(statArr[3])
					if err != nil {
						return err
					} else {
						if parentPid == pid {
							syscall.Kill(curPid, syscall.SIGKILL)
							break
						}
					}
				}
			}
		}
	}
	return syscall.Kill(pid, syscall.SIGKILL)
}

//判断进程是否存在
func exist(pid int) bool {
	ret := false
	err := syscall.Kill(pid, 0)
	if err == nil {
		ret = true
	}
	return ret
}
