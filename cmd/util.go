package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/irgendwr/go-stine/api"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

const keyUsername = "username"
const keyPassword = "password"

func credentials() (user string, pass string) {
	user = viper.GetString(keyUsername)
	pass = viper.GetString(keyPassword)

	if user == "" {
		fmt.Print("User: ")
		fmt.Scanln(&user)
	}

	if pass == "" {
		fmt.Print("Password: ")
		rpass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		pass = string(rpass)
		fmt.Print("\n")
	}

	return user, pass
}

// DownloadFile downloads a file from a given STiNE URL to disk.
func DownloadFile(acc api.Account, filepath string, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := acc.DoRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
