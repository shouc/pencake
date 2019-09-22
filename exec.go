package pencake

import "pencake/utils"

func Exec(command string) string {
	return utils.RunCommand(command)
}
