package pencake

import (
	"pencake/utils"
)

func DisableRemote() {
	switch utils.SystemType() {
	case "windows":
		utils.RunCommand(`REG ADD "HKLM\SYSTEM\CurrentControlSet\Control\Terminal Server" ` +
			`/v fDenyTSConnections /t REG_DWORD /d 1 /f`)
		return
	case "darwin":
		utils.RunCommand("sudo launchctl unload /System/Library/LaunchDaemons/ssh.plist")
		return
	case "linux":
		utils.RunCommand("service ssh stop")
		return
	default:
		utils.NotSupported("")
		return
	}

}
