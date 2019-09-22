package pencake

import (
	"fmt"
	"pencake/utils"
	"regexp"
)

func WifiKey() []WifiKeyStruct {
	switch utils.SystemType() {
	case "darwin":
		return getMacWifiKey()
	default:
		utils.NotSupported("")
		return []WifiKeyStruct{}
	}
}

type WifiKeyStruct struct {
	Name     string
	Password string
}

func getMacWifiKey() []WifiKeyStruct {
	wifiSSIDs := utils.RunCommand("defaults read " +
		"/Library/Preferences/SystemConfiguration/com.apple.airport.preferences KnownNetworks " +
		"| grep 'SSIDString'")
	fmt.Println(wifiSSIDs)
	wifiNameRegex := regexp.MustCompile(` {8}SSIDString = (.+?);`)
	wifiNames := wifiNameRegex.FindAllStringSubmatch(wifiSSIDs, -1)
	var result []WifiKeyStruct
	for _, name := range wifiNames {
		result = append(result, WifiKeyStruct{
			Name:     name[1],
			Password: utils.RunCommand(fmt.Sprintf("security find-generic-password -wa %s", name[1])),
		})
	}
	return result
}
