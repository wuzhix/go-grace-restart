package restart

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

const (
	procPath = "/proc"
)

func kill(pid int) error {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = windowsKill(pid)
	case "linux":
		err = linuxKill(pid)
	default:
	}
	return err
}

//windows下删除进程pid及其子进程
func windowsKill(pid int) error {
	process, err := os.FindProcess(pid)
	if err == nil {
		return process.Kill()
	}
	return err
}

//linux下删除进程pid及其子进程
func linuxKill(pid int) error {
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
	return nil
}

//判断进程是否存在
func exist(pid int) bool {
	ret := false
	switch runtime.GOOS {
	case "windows":
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
	case "linux":
		err := syscall.Kill(pid, 0)
		if err == nil {
			ret = true
		}
	default:
	}
	return ret
}
