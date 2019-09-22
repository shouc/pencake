package pencake

import (
	"pencake/utils"
)

func RemoteControl() {
	switch utils.SystemType() {
	case "windows":
		startRdp()
		return
	case "darwin":
		startSsh()
		return
	default:
		utils.NotSupported("")
		return
	}
}

func startRdp() {
	utils.RunCommand(`reg add "HKLM\SYSTEM\CurrentControlSet\Control\Terminal Server" ` +
		`/v fDenyTSConnections /t REG_DWORD /d 0 /f`)
}

func startSsh() {
	utils.RunCommand("systemsetup -setremotelogin on")
}
