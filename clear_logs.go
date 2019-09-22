package pencake

import (
	"pencake/utils"
	"strings"
)

func ClearLogs() bool {
	switch utils.SystemType() {
	case "windows":
		output1 := utils.RunCommand("wevtutil cl System")
		output2 := utils.RunCommand("wevtutil cl Security")
		output3 := utils.RunCommand("wevtutil cl Application")
		if strings.Contains(output1, "Access is denied") {
			return false
		}
		if strings.Contains(output2, "Access is denied") {
			return false
		}
		if strings.Contains(output3, "Access is denied") {
			return false
		}
		return true
	default:
		utils.NotSupported("")
		return false
	}
}
