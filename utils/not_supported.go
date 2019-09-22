package utils

import "fmt"

func NotSupported(sys string) {
	if sys == "" {
		sys = SystemType()
	}
	fmt.Println("[Error]: \033[1;37m" + sys + " is not supported for this function\033[0m")
}

func LinuxNotSupported() {
	NotSupported("linux")
}

func WindowsNotSupported() {
	NotSupported("windows")
}

func MacNotSupported() {
	NotSupported("mac")
}

func UnknownSystem() {
	NotSupported("unknown system")
}
