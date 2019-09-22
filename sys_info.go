package pencake

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"pencake/utils"
	"strings"
)

func SysInfo() SysInfoStruct {
	switch utils.SystemType() {
	case "windows":
		return getWindowsSysInfo()
	default:
		return getUnixSysInfo()
	}
}

type SysInfoStruct struct {
	Os       string
	Arch     string
	UserName string
	Name     string
	IsAdmin  bool
	PublicIp string
}

func getPublicIp() string {
	resp, _ := http.Get("https://ifconfig.me")
	publicIp, _ := ioutil.ReadAll(resp.Body)
	return string(publicIp)
}

func getUnixSysInfo() SysInfoStruct {
	var result SysInfoStruct
	userInfo, _ := user.Current()
	result.Os = utils.SystemType()
	result.Arch = utils.RunCommand("uname -m")
	result.UserName = userInfo.Username
	result.Name = userInfo.Name
	result.IsAdmin = os.Getuid() == 0
	result.PublicIp = getPublicIp()
	return result
}

func getWindowsSysInfo() SysInfoStruct {
	var result SysInfoStruct
	result.Os = "windows"
	archInfos := strings.Split(utils.RunCommand("wmic os get osarchitecture"), "\n")
	if len(archInfos) > 1 {
		result.Arch = archInfos[1]
	} else {
		result.Arch = "UNKNOWN"
	}
	userInfo, _ := user.Current()
	result.UserName = userInfo.Username
	result.Name = userInfo.Name
	userSID := userInfo.Gid
	if len(userSID) > 4 && userSID[len(userSID)-5:] == "-500" {
		result.IsAdmin = true
	}
	result.PublicIp = getPublicIp()
	return result
}
