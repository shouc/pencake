package pencake

import (
	"fmt"
	"pencake/utils"
)

func Popup(message string) {
	switch utils.SystemType() {
	case "darwin":
		utils.RunCommand(fmt.Sprintf(`osascript -e 'tell app "System Events" `+
			`to display dialog "%s" buttons {{"OK"}} default button "OK" '`, message))
		return
	case "windows":
		utils.RunCommand(fmt.Sprintf(`powershell "(new-object -ComObject wscript.shell)`+
			`.Popup(\"%s\",0,\"Windows\")"`, message))
		return
	default:
		utils.NotSupported("")
		return
	}
}
