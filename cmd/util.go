package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/irgendwr/go-stine/api"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

const keyUsername = "username"
const keyPassword = "password"
const keyNocache = "nocache"
const keyCacheTime = "cache_time"
const keyCacheSession = "cache_session"
const keyCacheCnsc = "cache_cnsc"

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

func login() (*api.Account, error) {
	acc := api.NewAccount()

	nocache := viper.GetBool(keyNocache)
	session, cnsc := viper.GetString(keyCacheSession), viper.GetString(keyCacheCnsc)
	if nocache || time.Since(viper.GetTime(keyCacheTime)) >= 30*time.Minute || session == "" || cnsc == "" {
		err := acc.Login(credentials())
		if !nocache {
			viper.Set(keyCacheTime, time.Now())
			session, cnsc = acc.Session()
			viper.Set(keyCacheSession, session)
			viper.Set(keyCacheCnsc, cnsc)
			viper.WriteConfig()
		}
		return &acc, err
	}
	acc.SetSession(session, cnsc)
	viper.Set(keyCacheTime, time.Now())
	viper.WriteConfig()
	return &acc, nil
}

// DownloadFile downloads a file from a given STiNE URL to disk.
func DownloadFile(acc *api.Account, filepath string, url string) error {
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
